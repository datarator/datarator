package fake

import (
	"math/rand"
	"strings"
)

// Character generates random character in the given language
func Character() string {
	return lookup(lang, "characters", true)
}

// CharactersN generates n random characters in the given language
func CharactersN(n int) string {
	var chars []string
	for i := 0; i < n; i++ {
		chars = append(chars, Character())
	}
	return strings.Join(chars, "")
}

// Characters generates from 1 to 5 characters in the given language
func Characters() string {
	return CharactersN(rand.Intn(5) + 1)
}

// Word generates random word
func Word() string {
	return lookup(lang, "words", true)
}

// WordsN generates n random words
func WordsN(n int) string {
	words := make([]string, n)
	for i := 0; i < n; i++ {
		words[i] = Word()
	}
	return strings.Join(words, " ")
}

func containsWord(words []string, word string) bool {
	for _, w := range words {
		if w == word {
			return true
		}
	}
	return false
}

// WordsNUnique generates n unique random words
func WordsNUnique(n int) []string {
	words := make([]string, n)
	word := Word()
	for i := 0; i < n; {
		if word != "" && !containsWord(words, word) {
			words[i] = word
			i++
			word = ""
		}
		word = word + Word()
	}
	return words
}

// Words generates from 1 to 5 random words
func Words() string {
	return WordsN(rand.Intn(5) + 1)
}

// Title generates from 2 to 5 titleized words
func Title() string {
	return strings.ToTitle(WordsN(2 + rand.Intn(4)))
}

// Sentence generates random sentence
func Sentence() string {
	var words []string
	for i := 0; i < 3+rand.Intn(12); i++ {
		word := Word()
		if rand.Intn(5) == 0 {
			word += ","
		}
		words = append(words, word)
	}

	sentence := strings.Join(words, " ")

	if rand.Intn(8) == 0 {
		sentence += "!"
	} else {
		sentence += "."
	}

	return sentence
}

// SentencesN generates n random sentences
func SentencesN(n int) string {
	sentences := make([]string, n)
	for i := 0; i < n; i++ {
		sentences[i] = Sentence()
	}
	return strings.Join(sentences, " ")
}

// Sentences generates from 1 to 5 random sentences
func Sentences() string {
	return SentencesN(rand.Intn(5) + 1)
}

// Paragraph generates paragraph
func Paragraph() string {
	return SentencesN(rand.Intn(10) + 1)
}

// ParagraphsN generates n paragraphs
func ParagraphsN(n int) string {
	var paragraphs []string
	for i := 0; i < n; i++ {
		paragraphs = append(paragraphs, Paragraph())
	}
	return strings.Join(paragraphs, "\t")
}

// Paragraphs generates from 1 to 5 paragraphs
func Paragraphs() string {
	return ParagraphsN(rand.Intn(5) + 1)
}
