package textutil

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
	// NegationScalar empirically derived mean sentiment intensity reflection upon negating
	NegationScalar float64 = -0.74
)

// SentimentIdioms using a sentiment-laden laden idioms that do not contain lexicon words (future work, not yet implemented)
var SentimentIdioms map[string]float64 = map[string]float64{"cut the mustard": 2, "hand to mouth": -2,
	"back handed": -2, "blow smoke": -2, "blowing smoke": -2,
	"upper hand": 1, "break a leg": 2,
	"cooking with gas": 2, "in the black": 2, "in the red": -2,
	"on the ball": 2, "under the weather": -2}

// SpecialIdioms check for special case idioms containing lexicon words
var SpecialIdioms map[string]float64 = map[string]float64{"the shit": 3, "the bomb": 3, "bad ass": 1.5, "badass": 1.5,
	"yeah right": -2, "kiss of death": -1.5, "to die for": 3}
