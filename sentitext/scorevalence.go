package sentitext

import (
	"math"
	"strings"

	textutil "github.com/grassmudhorses/vader-go/internal/textutil"
)

// Sentiment of a phrase or sentence
type Sentiment struct {
	Negative float64 `json:"neg" csv:"neg"`
	Neutral  float64 `json:"neu" csv:"neu"`
	Positive float64 `json:"pos" csv:"pos"`
	Compound float64 `json:"compound" csv:"compound"`
}

// ScoreValence .
func ScoreValence(sentimentscores []float64, text string) Sentiment {
	if len(sentimentscores) != 0 {
		var sentimentSum float64
		for _, v := range sentimentscores {
			sentimentSum = sentimentSum + float64(v)
		}

		// compute and add emphasis from punctuation in text
		puncAmp := punctuationEmphasis(text)

		if sentimentSum > 0 {
			sentimentSum += puncAmp
		} else if sentimentSum < 0 {
			sentimentSum -= puncAmp
		}
		compound := normalize(sentimentSum, 15.0)
		// discriminate between positive, negative and neutral sentiment scores
		pos, neg, neu := siftSentimentScores(sentimentscores)
		if pos > math.Abs(neg) {
			pos += puncAmp
		} else if pos < math.Abs(neg) {
			neg -= puncAmp
		}

		total := pos + math.Abs(neg) + neu

		return Sentiment{
			Negative: math.Abs(neg / total),
			Positive: math.Abs(pos / total),
			Neutral:  math.Abs(neu / total),
			Compound: compound,
		}
	}
	return Sentiment{}
}

// punctuationEmphasis adds emphasis from exclamation points and question marks
func punctuationEmphasis(text string) float64 {
	return amplifyExclaim(text) + amplifyQuestion(text)
}

// check for added emphasis resulting from exclamation points (up to 4 of them)
func amplifyExclaim(text string) float64 {
	epCount := strings.Count(text, "!")
	if epCount > 4 {
		return 4.0 * textutil.BoosterIncrease
	}
	// (empirically derived mean sentiment intensity rating increase for exclamation points)
	return float64(epCount) * textutil.BoosterIncrease
}

// check for added emphasis resulting from question marks (2 or 3+)
func amplifyQuestion(text string) float64 {
	qmCount := float64(strings.Count(text, "?"))
	switch {
	// (empirically derived mean sentiment intensity rating increase for question marks)
	case qmCount <= 1:
		return 0
	case qmCount <= 3:
		return qmCount * 0.18
	default:
		return 0.96
	}
}

func siftSentimentScores(sentimentscores []float64) (posSum float64, negSum float64, neuCount float64) {
	var neu int
	for _, score := range sentimentscores {
		switch {
		case score > 0:
			posSum += score + 1.0 //compensates for neutral words that are counted as 1
		case score < 0:
			negSum += score - 1.0 //compensates for neutrals
		case score == 0:
			neu++
		}
	}
	return posSum, negSum, float64(neu)
}

//normalize the score to be between -1 and 1 using an alpha that approximates the max expected value
func normalize(score float64, alpha float64) float64 {
	norm := score / math.Sqrt((score*score)+alpha)
	switch {
	case norm < -1.0:
		return -1.0
	case norm > 1.0:
		return 1.0
	default:
		return norm
	}
}
