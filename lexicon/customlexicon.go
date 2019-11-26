package lexicon

//CustomLexicon can be used to create your own lexicon maps, and be used for mocking
type CustomLexicon struct {
	//This just acts as the default lexicon based on static files
	NegateList map[string]bool
	Contrasts  map[string]bool
	Sentiments map[string]float64
	Boosters   map[string]float64
}

func (s *CustomLexicon) IsNegation(text string) bool {
	return s.NegateList[text]
}
func (s *CustomLexicon) IsContrast(text string) bool {
	return s.Contrasts[text]
}
func (s *CustomLexicon) Sentiment(text string) float64 {
	return s.Sentiments[text]
}
func (s *CustomLexicon) BoostValue(text string) float64 {
	return s.Boosters[text]
}
