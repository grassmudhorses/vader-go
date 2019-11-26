package sentitext

import (
	"regexp"
	"strings"
	"unicode"

	"github.com/grassmudhorses/vader-go/lexicon"
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
	WordsAndEmotes *[]SentiWord
	IsCapDiff      bool
	Original       string
}

type SentiWord struct {
	Word          string
	Lower         string
	BaseSentiment float64
	IsCaps        bool
	BoostValue    float64
	IsContrast    bool
	IsNegation    bool
}

// Parse and Identify sentiment-relevant string-level properties of input text
func Parse(text string, lex lexicon.Lexicon) (s *SentiText) {
	s = &SentiText{}
	sentwords := getWordsAndEmoticons(text, lex)
	s.WordsAndEmotes = &sentwords
	s.IsCapDiff = allCapsDifferential(s.WordsAndEmotes)
	s.Original = text
	return
}

// getWordsAndEmoticons Removes leading and trailing puncutation Leaves contractions
// and emoji. Does not preserve punc-plus-letter emoticons (e.g. :D)
func getWordsAndEmoticons(text string, lex lexicon.Lexicon) []SentiWord {
	wordsOnly := []SentiWord{}
	for _, token := range Spaces.FindAllString(text, -1) {
		for _, word := range splitEmojis(token) {
			word = strings.TrimFunc(word, unicode.IsPunct)
			if len(word) != 0 {
				lower := strings.ToLower(word)
				isUpper := strings.ToUpper(word) == word && !Emoji.MatchString(word)
				wordsOnly = append(wordsOnly, SentiWord{
					BaseSentiment: lex.Sentiment(lower),
					IsCaps:        isUpper,
					Lower:         lower,
					Word:          word,
					BoostValue:    lex.BoostValue(word),
					IsContrast:    lex.IsContrast(word),
					IsNegation:    lex.IsNegation(word),
				})
			}
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
func allCapsDifferential(words *[]SentiWord) bool {
	var totallength int
	var capslength int
	for _, word := range *words {
		//emojis and single letters ignored for caps
		if len(word.Word) < 2 || Emoji.MatchString(word.Word) {
			continue
		}
		totallength += len(word.Word)
		if word.IsCaps {
			capslength += len(word.Word)
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
