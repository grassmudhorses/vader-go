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
	// NScalar .
	NScalar float64 = -0.74
)
