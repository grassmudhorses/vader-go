package textutil

import (
	"math"
	"regexp"
	"strings"
)

// Go port of vader sentiment analysis tool, source:
// Hutto, C.J. & Gilbert, E.E. (2014). VADER`: A Parsimonious Rule-based Model for
// Sentiment Analysis of Social Media Text. Eighth International Conference on
// Weblogs and Social Media (ICWSM-14). Ann Arbor, MI, June 2014.

const (
	// BoosterIncrease empirically derived mean sentiment intensity rating increase for booster words
	BoosterIncrease float64 = 0.293
	// BoosterDecrease empirically derived mean sentiment intensity rating increase for booster words
	BoosterDecrease float64 = -0.293
	// CapsIncrease empirically derived mean sentiment intensity rating increase for using ALLCAPs to emphasize a word
	CapsIncrease float64 = 0.733
	// NScalar .
	NScalar float64 = -0.74
)

// NonWords simple regex to split english
var NonWords *regexp.Regexp

func init() {
	NonWords = regexp.MustCompile(`[\'\w\d]+`)
}

// AllCapsDifferential Check whether just some words in the input are ALL CAPS
func AllCapsDifferential(words []string) bool {
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

//Normalize the score to be between -1 and 1 using an alpha that approximates the max expected value
func Normalize(score float64, alpha float64) float64 {
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
