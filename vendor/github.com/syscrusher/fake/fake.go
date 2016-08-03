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

// cat/subcat/lang/samples
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

var r = rand.New(&rndSrc{src: rand.NewSource(time.Now().UnixNano())})
var lang = "en"
var useExternalData = false
var enFallback = true
var availLangs = GetLangs()

var (
	// ErrNoLanguage indicates that the given language is not available
	ErrNoLanguage = errors.New("Language unavailable")
	// ErrNoSamples indicates that there are no samples for the given language
	ErrNoSamples = errors.New("No samples found for given language")
)

type rndSrc struct {
	mtx sync.Mutex
	src rand.Source
}

func (s *rndSrc) Int63() int64 {
	s.mtx.Lock()
	n := s.src.Int63()
	s.mtx.Unlock()
	return n
}

func (s *rndSrc) Seed(n int64) {
	s.mtx.Lock()
	s.src.Seed(n)
	s.mtx.Unlock()
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

// SetLang sets the language in which the data should be generated
// returns error if passed language is not available
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

// UseExternalData sets the flag that allows using of external files as data providers (fake uses embedded ones by default)
func UseExternalData(flag bool) {
	useExternalData = flag
}

// EnFallback sets the flag that allows fake to fallback to englsh samples if the ones for the used languaged are not available
func EnFallback(flag bool) {
	enFallback = flag
}

func join(parts ...string) string {
	var filtered []string
	for _, part := range parts {
		if part != "" {
			filtered = append(filtered, part)
		}
	}
	return strings.Join(filtered, " ")
}

func generate(lag, cat string, fallback bool) string {
	format := lookup(lang, cat+"_format", fallback)
	var result string
	for _, ru := range format {
		if ru != '#' {
			result += string(ru)
		} else {
			result += strconv.Itoa(r.Intn(10))
		}
	}
	return result
}

func lookup(lang, cat string, fallback bool) string {
	var s []string

	samples.RLock()
	if samples.cache.hasKeyPath(lang, cat) {
		s = samples.cache[lang][cat]
		samples.RUnlock()
		return s[r.Intn(len(s))]
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
	return s[r.Intn(len(s))]
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
