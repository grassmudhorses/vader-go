package sentitext

import (
	"github.com/grassmudhorses/vader-go/internal/lexicon"
)

// SentiText Identify sentiment-relevant string-level properties of input text
type SentiText struct {
	Text             string
	WordsAndEmotions []string
	IsCapDiff        bool
}

// Parse and Identify sentiment-relevant string-level properties of input text
func Parse(text string) (s *SentiText) {
	s.Text = text
	//   this.words_and_emoticons = this.get_words_and_emoticons();
	//   // doesn't separate words from adjacent punctuation (keeps emoticons & contractions)
	//   this.is_cap_diff = allcap_differential(this.words_and_emoticons);
	s.WordsAndEmotions = getWordsAndEmoticons(text)
	s.IsCapDiff = lexicon.AllCapsDifferential(s.WordsAndEmotions)
	return
}

// getWordsAndEmoticons Removes leading and trailing puncutation Leaves contractions and most emoticons Does not preserve punc-plus-letter emoticons (e.g. :D)
func getWordsAndEmoticons(text string) []string {
	wordsPuncDict := getWordsPlusPunc(text)
	wordsOnly := []string{}
	for _, token := range lexicon.Spaces.FindAllString(text, -1) {
		if len(token) < 2 {
			continue
		}
		wordsOnly = append(wordsOnly, wordsPuncDict[token])
	}
	return wordsOnly
}

// getWordsPlusPunc Returns mapping of form:
// {  'cat,': 'cat',  ',cat': 'cat'}
func getWordsPlusPunc(text string) map[string]string {
	noPuncText := lexicon.Punc.ReplaceAllLiteralString(text, "")
	wordsPuncDict := make(map[string]string)
	for _, word := range lexicon.Spaces.FindAllString(noPuncText, -1) {
		if len(word) < 2 {
			continue
		}
		for _, punc := range lexicon.PunctuationList {
			wordsPuncDict[word+punc] = word
			wordsPuncDict[punc+word] = word
		}
	}
	return wordsPuncDict
}
