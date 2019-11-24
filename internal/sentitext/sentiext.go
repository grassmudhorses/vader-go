package sentitext

import (
	"regexp"
	"strings"
	"unicode"
)

// Spaces simple regex to split emojis and words
var Spaces *regexp.Regexp

// Emoji simple regex to match only emoji
var Emoji *regexp.Regexp

func init() {
	Spaces = regexp.MustCompile(`[^\p{Z}\p{C}\p{Sm}\p{Sc}]+`)
	Emoji = regexp.MustCompile(`[\p{So}\p{Sk}]+`)
}

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
// and emoji. Does not preserve punc-plus-letter emoticons (e.g. :D)
func getWordsAndEmoticons(text string) []string {
	wordsOnly := []string{}
	for _, token := range Spaces.FindAllString(text, -1) {
		for _, word := range splitEmojis(token) {
			word = strings.TrimFunc(word, unicode.IsPunct)
			wordsOnly = append(wordsOnly, word)
		}
	}
	return wordsOnly
}

// splitEmojis out of a word with attached emojis
func splitEmojis(token string) []string {
	out := []string{}
	var loc []int
	for i := 0; i < len(token); {
		loc = Emoji.FindStringIndex(token[i:len(token)])
		//find the next emoji at  token[i+loc[0]:i+loc[1]]
		if loc != nil {
			//if there's words before the emoji
			if loc[0] > 0 {
				out = append(out, token[i:i+loc[0]])
				out = append(out, token[i+loc[0]:i+loc[1]])
			} else {
				out = append(out, token[i+loc[0]:i+loc[1]])
			}
			i += loc[1]
		} else {
			out = append(out, token[i:len(token)])
			break
		}
	}
	return out
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
