package lexicon

type Lexicon interface {
	// Negation words like aren't, ain't, wouldn't
	IsNegation(string) bool
	// Contrasting conjuctions like however, but, yet
	IsContrast(string) bool
	// Base sentiment value of a word, >0 is positive sentiment, <0 is negative sentiment, =0 no sentiment
	Sentiment(string) float64
	// value of a word, >0 is positive sentiment, <0 is negative sentiment, =0 no sentiment
	BoostValue(string) float64
	//TODO: Idioms
}

type StaticLexicon struct {
	//This just acts as the default lexicon based on static files
}

func (s StaticLexicon) IsNegation(text string) bool {
	return NegateList[text]
}
func (s StaticLexicon) IsContrast(text string) bool {
	return Contrasts[text]
}
func (s StaticLexicon) Sentiment(text string) float64 {
	return Sentiments[text]
}
func (s StaticLexicon) BoostValue(text string) float64 {
	return Boosters[text]
}

//DefaultLexicon to look up on the static list provided by VADER
var DefaultLexicon = StaticLexicon{}
