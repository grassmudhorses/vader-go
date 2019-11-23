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
	s.IsCapDiff = allCapsDifferential(s.WordsAndEmotes)
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

// allCapsDifferential Check whether just some words in the input are ALL CAPS
func allCapsDifferential(words []string) bool {
	var totallength int
	var capslength int
	for _, word := range words {
		totallength += len(word)
		if strings.ToUpper(word) == word {
			capslength += len(word)
		}
	}
	//only true if words are partially caps, and at least 10% of letters are caps
	if capslength == 0 || totallength == capslength {
		return false
	}
	if capslength > 20 || float32(capslength)*10.0 > float32(totallength) {
		return true
	}
	return false
}
