package sentitext

import (
	"strings"

	textutil "github.com/grassmudhorses/vader-go/internal/textutil"
)

// PolarityScore Return a float for sentiment strength based on the input text. Positive values are positive valence, negative value are negative valence.
func PolarityScore(text string) Sentiment {
	senti := Parse(text)
	sentiments := []float64{}
	for i, word := range senti.WordsAndEmotes {
		//special case for "kind of"
		if i < len(senti.WordsAndEmotes)-1 && strings.ToLower(word) == "kind" && strings.ToLower(senti.WordsAndEmotes[i+1]) == "of" {
			sentiments = append(sentiments, 0.0)
			continue
		}
		//boost words don't have valence since they can be positive or negative
		if textutil.Lexicon[strings.ToLower(word)] != 0.0 {
			sentiments = append(sentiments, 0.0)
			continue
		}
		//determine sentiment of current word based on surrounding lexicon cues
		sentiments = SentimentValence(senti, word, i, sentiments)
	}
	//apply "but" check
	sentiments = butCheck(senti.WordsAndEmotes, sentiments)
	return ScoreValence(sentiments, text)
}

// butCheck check for modification in sentiment due to contrastive conjunction 'but'
func butCheck(words []string, sentiments []float64) []float64 {
	//FIXME: can be optimized to O(words+sentiments) by compling a list of all 'but' indicies
	for i, word := range words {
		if strings.ToLower(word) != "but" {
			continue
		}
		// every but makes all sentiments before it weaker, and those after it stronger.
		for j, sentiment := range sentiments {
			if j < i {
				sentiments[j] = sentiment * 0.5
			} else {
				sentiments[j] = sentiment * 1.5
			}
		}
	}
	return sentiments
}
