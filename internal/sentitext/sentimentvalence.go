package sentitext

// sentiment_valence(self, valence, sentitext, item, i, sentiments):
// 	is_cap_diff = sentitext.is_cap_diff
// 	words_and_emoticons = sentitext.words_and_emoticons
// 	item_lowercase = item.lower()
// 	if item_lowercase in self.lexicon:
// 		# get the sentiment valence
// 		valence = self.lexicon[item_lowercase]

// 		# check for "no" as negation for an adjacent lexicon item vs "no" as its own stand-alone lexicon item
// 		if item_lowercase == "no" and words_and_emoticons[i + 1].lower() in self.lexicon:
// 			# don't use valence of "no" as a lexicon item. Instead set it's valence to 0.0 and negate the next item
// 			valence = 0.0
// 		if (i > 0 and words_and_emoticons[i - 1].lower() == "no") \
// 			or (i > 1 and words_and_emoticons[i - 2].lower() == "no") \
// 			or (i > 2 and words_and_emoticons[i - 3].lower() == "no" and words_and_emoticons[i - 1].lower() in ["or", "nor"] ):
// 			valence = self.lexicon[item_lowercase] * N_SCALAR

// 		# check if sentiment laden word is in ALL CAPS (while others aren't)
// 		if item.isupper() and is_cap_diff:
// 			if valence > 0:
// 				valence += C_INCR
// 			else:
// 				valence -= C_INCR

// 		for start_i in range(0, 3):
// 			# dampen the scalar modifier of preceding words and emoticons
// 			# (excluding the ones that immediately preceed the item) based
// 			# on their distance from the current item.
// 			if i > start_i and words_and_emoticons[i - (start_i + 1)].lower() not in self.lexicon:
// 				s = scalar_inc_dec(words_and_emoticons[i - (start_i + 1)], valence, is_cap_diff)
// 				if start_i == 1 and s != 0:
// 					s = s * 0.95
// 				if start_i == 2 and s != 0:
// 					s = s * 0.9
// 				valence = valence + s
// 				valence = self._negation_check(valence, words_and_emoticons, start_i, i)
// 				if start_i == 2:
// 					valence = self._special_idioms_check(valence, words_and_emoticons, i)

// 		valence = self._least_check(valence, words_and_emoticons, i)
// 	sentiments.append(valence)
// 	return sentiments

func SentimentValence(sentitext *SentiText, item string, i int, sentiments []float64) []float64 {
	valence := 0.0
	return []float64{valence}
}
