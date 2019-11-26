package sentitext

import (
	"strings"

	"github.com/grassmudhorses/vader-go/textutil"
)

// SentimentValence .
func SentimentValence(sentitext *SentiText, item SentiWord, i int, sentiments []float64) []float64 {
	valence := item.BaseSentiment
	words := *sentitext.WordsAndEmotes
	if valence != 0.0 {
		//check for "no" as negation for an adjacent lexicon item vs "no" as its own stand-alone lexicon item
		if item.Lower == "no" {
			if i+1 < len(*sentitext.WordsAndEmotes) && words[i+1].BaseSentiment != 0 {
				//don't use valence of "no" as a lexicon item. Instead set it's valence to 0.0 and negate the next item
				valence = 0.0
			}
		}
		if (i > 0 && words[i-1].Lower == "no") || (i > 1 && words[i-2].Lower == "no") || (i > 2 && words[i-3].Lower == "no" && (words[i-1].Lower == "or" || words[i-1].Lower == "nor")) {
			valence = valence * textutil.NegationScalar
		}
		//check if sentiment laden word is in ALL CAPS (while others aren't)
		if item.IsCaps && sentitext.IsCapDiff {
			if valence > 0 {
				valence += textutil.CapsIncrease
			} else if valence < 0 {
				valence -= textutil.CapsIncrease
			}
		}
		var s float64
		//dampen the scalar modifier of preceding words and emoticons (excluding the ones that immediately preceed the item) based on their distance from the current item.
		for start := 0; start < 4; start++ {
			if i > start && words[i-(start+1)].BaseSentiment == 0.0 {
				s = scalarIncreaseDecease(words[i-(start+1)], valence, sentitext.IsCapDiff)
				switch start {
				case 1:
					s = s * 0.95
				case 2:
					s = s * 0.9
				}
				valence += s
				valence = negationCheck(valence, &words, start, i)
				if start == 2 {
					valence = specialIdiomsCheck(valence, &words, i)
				}
			}
		}
	}
	sentiments = append(sentiments, valence)
	return sentiments
}

// scalarIncreaseDecease Check if the preceding words increase, decrease, or negate/nullify the valence
func scalarIncreaseDecease(word SentiWord, valence float64, isCapDiff bool) float64 {
	scalar := word.BoostValue
	if scalar != 0.0 {
		if valence < 0 {
			scalar *= -1.0
		}
		//check if booster/dampener word is in ALLCAPS (while others aren't)
		if word.IsCaps && isCapDiff {
			if valence > 0 {
				scalar += textutil.CapsIncrease
			} else {
				scalar -= textutil.CapsIncrease
			}
		}
	}
	return scalar
}
func negationCheck(valence float64, wordsp *[]SentiWord, start int, i int) float64 {
	words := *wordsp
	switch start {
	case 0:
		if containsNegation(words[i-(start+1):]) {
			valence *= textutil.NegationScalar
		}
	case 1:
		if words[i-2].Lower == "never" && (words[i-1].Lower == "so" || words[i-1].Lower == "this") {
			valence *= 1.25
		} else if words[i-2].Lower == "without" && words[i-1].Lower == "doubt" {
			//do nothing
		} else if containsNegation(words[i-(start+1):]) {
			//2 words preceding the lexicon word position
			valence = valence * textutil.NegationScalar
		}
	case 2:
		if words[i-3].Lower == "never" && (words[i-2].Lower == "so" || words[i-2].Lower == "this") && (words[i-1].Lower == "so" || words[i-1].Lower == "this") {
			valence *= 1.25
		} else if words[i-3].Lower == "without" && (words[i-2].Lower == "doubt" || words[i-1].Lower == "doubt") {
			//do nothing
		} else if containsNegation(words[i-(start+1):]) {
			//3 words preceding the lexicon word position
			valence *= textutil.NegationScalar
		}

	}
	return valence
}

//leastCheck for negation case using "least"
func leastCheck(valence float64, wordsp *[]SentiWord, i int) float64 {
	words := *wordsp
	if i > 1 && words[i-1].BaseSentiment == 0 && words[i-1].Lower == "least" {
		if i > 0 || (words[i-2].Lower != "at" && words[i-2].Lower != "very") {
			valence = valence * textutil.NegationScalar
		}
	}
	return valence
}
func specialIdiomsCheck(valence float64, words *[]SentiWord, i int) float64 {
	return valence
}

// Determine if input contains negation words
func containsNegation(lowerwords []SentiWord) bool {
	for _, word := range lowerwords {
		if word.IsNegation {
			return true
		}
		if strings.Contains(word.Lower, "n't") {
			return true
		}
	}
	return false
}
