package fake

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

//go:generate go get github.com/mjibson/esc
//go:generate esc -o data.go -pkg fake data

func init() {
	// seed math/rand's global source so output is truly random
	rand.Seed(time.Now().UnixNano())
}

/*
	Language handling
*/

var (
	lang       = "en"
	enFallback = true
	availLangs = GetLangs()

	// ErrNoLanguage indicates that the given language is not available
	ErrNoLanguage = errors.New("Language unavailable")
	// ErrNoSamples indicates that there are no samples for the given language
	ErrNoSamples = errors.New("No samples found for given language")
)

// EnFallback sets the flag that allows fake to fallback to englsh samples if the ones for the used languaged are not available
func EnFallback(flag bool) {
	enFallback = flag
}

// GetLangs returns a slice of available languages
func GetLangs() []string {
	var langs []string
	for k, v := range _escData {
		if v.isDir && k != "/" && k != "/data" {
			langs = append(langs, strings.Replace(k, "/data/", "", 1))
		}
	}
	return langs
}

// SetLang sets the language in which the data should be generated,
// and returns ErrNoLanguage if lang doesn't exist
func SetLang(newLang string) error {
	found := false
	for _, l := range availLangs {
		if newLang == l {
			found = true
			break
		}
	}
	if !found {
		return ErrNoLanguage
	}
	lang = newLang
	return nil
}

/*
	Data files, caching, lookup, generation
*/

var useExternalData = false

// UseExternalData sets the flag that allows using of external files as data providers (fake uses embedded ones by default)
func UseExternalData(flag bool) {
	useExternalData = flag
}

// map[lang]map[cat][]samples
type cache map[string]map[string][]string

func (c cache) hasKeyPath(lang, cat string) bool {
	if _, ok := c[lang]; ok {
		if _, ok = c[lang][cat]; ok {
			return true
		}
	}
	return false
}

var samples = struct {
	sync.RWMutex
	cache
}{cache: make(cache)}

func lookup(lang, cat string, fallback bool) string {
	var s []string

	samples.RLock()
	if samples.cache.hasKeyPath(lang, cat) {
		s = samples.cache[lang][cat]
		samples.RUnlock()
		return s[rand.Intn(len(s))]
	}
	samples.RUnlock()

	var err error
	s, err = populateSamples(lang, cat)
	if err != nil {
		if lang != "en" && fallback && enFallback && err.Error() == ErrNoSamples.Error() {
			return lookup("en", cat, false)
		}
		return ""
	}
	return s[rand.Intn(len(s))]
}

func populateSamples(lang, cat string) ([]string, error) {
	data, err := readFile(lang, cat)
	if err != nil {
		return nil, err
	}

	s := strings.Split(strings.TrimSpace(string(data)), "\n")

	samples.Lock()
	if _, ok := samples.cache[lang]; !ok {
		samples.cache[lang] = make(map[string][]string)
	}
	samples.cache[lang][cat] = s
	samples.Unlock()

	return s, nil
}

func readFile(lang, cat string) (f []byte, err error) {
	fullpath := fmt.Sprintf("/data/%s/%s", lang, cat)
	file, err := FS(useExternalData).Open(fullpath)
	if err != nil {
		return nil, ErrNoSamples
	}
	defer func() {
		// overwrites named return value
		err = file.Close()
	}()

	return ioutil.ReadAll(file)
}

// generate reads a random line from cat_format for the given lang,
// then replaces all instances of '#' with a random int 0-9 inclusive.
func generate(lang, cat string, fallback bool) string {
	format := lookup(lang, cat+"_format", fallback)
	var result string
	for _, ru := range format {
		if ru != '#' {
			result += string(ru)
		} else {
			result += strconv.Itoa(rand.Intn(10))
		}
	}
	return result
}

/*
	Helpers
*/

// joins non-blank string parts delimited by spaces
func join(parts ...string) string {
	var filtered []string
	for _, part := range parts {
		if part != "" {
			filtered = append(filtered, part)
		}
	}
	return strings.Join(filtered, " ")
}
