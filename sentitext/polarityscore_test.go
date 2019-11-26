package sentitext

import (
	"reflect"
	"testing"

	"github.com/grassmudhorses/vader-go/lexicon"
)

func TestPolarityScore(t *testing.T) {
	tests := []struct {
		name string
		text string
		want Sentiment
	}{
		{
			name: "navyseal1",
			text: `I'll have you know I graduated top of my class in the Navy Seals, and I've been involved in numerous secret raids on Al-Quaeda, and I have over 300 confirmed kills.`,
			want: Sentiment{
				Negative: 0.10204081632653061,
				Neutral:  0.8454810495626823,
				Positive: 0.05247813411078718,
				Compound: -0.401923825269382,
			},
		},
		{
			name: "navyseal2",
			text: `I will wipe you the fuck out with precision the likes of which has never been seen before on this Earth, mark my fucking words. `,
			want: Sentiment{
				Negative: 0.0646688073221318,
				Neutral:  0.8600554043204218,
				Positive: 0.07527578835744646,
				Compound: 0.07304422535683062,
			},
		},
		{
			name: "navyseal3",
			text: `As we speak I am contacting my secret network of spies across the USA and your IP is being traced right now so you better prepare for the storm, maggot. `,
			want: Sentiment{
				Negative: 0.0,
				Neutral:  0.9012270672672775,
				Positive: 0.09877293273272246,
				Compound: 0.4902265129795313,
			},
		},
		{
			name: "navyseal4",
			text: `Not only am I extensively trained in unarmed combat, but I have access to the entire arsenal of the United States Marine Corps and I will use it to its full extent to wipe your miserable ass off the face of the continent, you little shit. `,
			want: Sentiment{
				Negative: 0.25388704818020214,
				Neutral:  0.6843541616244231,
				Positive: 0.06175879019537477,
				Compound: -0.9101820964684066,
			},
		},
		{
			name: "navyseal5",
			text: `If only you could have known what unholy retribution your little "clever" comment was about to bring down upon you, maybe you would have held your fucking tongue. `,
			want: Sentiment{
				Negative: 0.0,
				Neutral:  0.9088766957282796,
				Positive: 0.09112330427172047,
				Compound: 0.40331007072320413,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sentis := Parse(tt.text, lexicon.DefaultLexicon)
			if got := PolarityScore(sentis); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PolarityScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
