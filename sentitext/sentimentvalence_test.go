package sentitext

import "testing"

func Test_containsNegation(t *testing.T) {
	tests := []struct {
		name       string
		lowerwords []SentiWord
		want       bool
	}{
		{
			name: "contains a don't",
			lowerwords: []SentiWord{
				{
					Lower: "contains a don't",
				},
			},
			want: true,
		},
		{
			name: "contains a wouldn't",
			lowerwords: []SentiWord{
				{
					Lower: "contains a wouldn't",
				},
			},
			want: true,
		},

		{
			name: "contains nothing",
			lowerwords: []SentiWord{
				{
					Lower: "contains nothing",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := containsNegation(tt.lowerwords); got != tt.want {
				t.Errorf("containsNegation() = %v, want %v", got, tt.want)
			}
		})
	}
}
