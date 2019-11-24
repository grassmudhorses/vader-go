package sentitext

import (
	textutil "github.com/grassmudhorses/vader-go/internal/textutil"
)

// PolarityScore Return a float for sentiment strength based on the input text. Positive values are positive valence, negative value are negative valence.
func PolarityScore(text string) Sentiment {
	senti := Parse(text)

	words := *senti.WordsAndEmotes
	sentiments := []float64{}
	for i, word := range words {
		//special case for "kind of"
		if i < len(*senti.WordsAndEmotes)-1 && word.Lower == "kind" && words[i+1].Lower == "of" {
			sentiments = append(sentiments, 0.0)
			continue
		}
		//boost words don't have valence since they can be positive or negative
		if textutil.Boosters[word.Lower] != 0.0 {
			sentiments = append(sentiments, 0.0)
			continue
		}
		//determine sentiment of current word based on surrounding lexicon cues
		sentiments = SentimentValence(senti, word, i, sentiments)
	}
	//apply "but" check
	sentiments = butCheck(&words, sentiments)
	return ScoreValence(sentiments, text)
}

// butCheck check for modification in sentiment due to contrastive conjunction 'but'
func butCheck(words *[]SentiWord, sentiments []float64) []float64 {
	//FIXME: can be optimized to O(words+sentiments) by compling a list of all 'but' indicies
	for i, word := range *words {
		if word.Lower != "but" {
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
