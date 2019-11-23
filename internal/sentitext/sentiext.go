package sentitext

import textutil "github.com/grassmudhorses/vader-go/internal/textutil"

// SentiText Identify sentiment-relevant string-level properties of input text
type SentiText struct {
	WordsAndEmotes []string
	IsCapDiff      bool
}

// Parse and Identify sentiment-relevant string-level properties of input text
func Parse(text string) (s *SentiText) {
	s = &SentiText{}
	s.WordsAndEmotes = getWordsAndEmoticons(textutil.CleanExtraPunc(text))
	//TODO: clean out repeat characters
	s.IsCapDiff = textutil.AllCapsDifferential(s.WordsAndEmotes)
	return
}

// getWordsAndEmoticons Removes leading and trailing puncutation Leaves contractions and most emoticons Does not preserve punc-plus-letter emoticons (e.g. :D)
func getWordsAndEmoticons(text string) []string {
	wordsPuncDict := getWordsPlusPunc(text)
	wordsOnly := []string{}
	for _, token := range textutil.Spaces.FindAllString(text, -1) {
		if len(token) < 2 {
			continue
		}
		word := wordsPuncDict[token]
		if word != "" {
			wordsOnly = append(wordsOnly, word)
		} else {
			wordsOnly = append(wordsOnly, token)
		}
	}
	return wordsOnly
}

// getWordsPlusPunc Returns mapping of form:
// {  'cat,': 'cat',  ',cat': 'cat'}
func getWordsPlusPunc(text string) map[string]string {
	noPuncText := textutil.Punc.ReplaceAllLiteralString(text, "")
	wordsPuncDict := make(map[string]string)
	for _, word := range textutil.Spaces.FindAllString(noPuncText, -1) {
		if len(word) < 2 {
			continue
		}
		for _, punc := range textutil.PunctuationList {
			wordsPuncDict[word+punc] = word
			wordsPuncDict[punc+word] = word
		}
	}
	return wordsPuncDict
}
