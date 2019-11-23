package sentitext

import (
	"strings"
	"unicode"

	textutil "github.com/grassmudhorses/vader-go/internal/textutil"
)

// SentiText Identify sentiment-relevant string-level properties of input text
type SentiText struct {
	WordsAndEmotes []string
	IsCapDiff      bool
}

// Parse and Identify sentiment-relevant string-level properties of input text
func Parse(text string) (s *SentiText) {
	s = &SentiText{}
	s.WordsAndEmotes = getWordsAndEmoticons(text)
	//TODO: clean out repeat characters
	s.IsCapDiff = textutil.AllCapsDifferential(s.WordsAndEmotes)
	return
}

// getWordsAndEmoticons Removes leading and trailing puncutation Leaves contractions
// and most emoticons Does not preserve punc-plus-letter emoticons (e.g. :D)
func getWordsAndEmoticons(text string) []string {
	wordsOnly := []string{}
	for _, token := range textutil.NonWords.FindAllString(text, -1) {
		cutToken := strings.TrimFunc(token, unicode.IsPunct)
		if len(cutToken) > 2 {
			wordsOnly = append(wordsOnly, cutToken)
		}
	}
	return wordsOnly
}
