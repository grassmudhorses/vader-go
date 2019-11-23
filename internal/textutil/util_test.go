package textutil

import (
	"testing"
)

func TestAllCapsDifferential(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  bool
	}{
		{
			name:  "allcaps",
			words: []string{"WHAT'S", "THE", "GOOD", "WORD?"},
			want:  false,
		},
		{
			name:  "alllower",
			words: []string{"To", "Hell", "With", "georgia!"},
			want:  false,
		},
		{
			name:  "morecaps",
			words: []string{"WHAT'S", "the", "GOOD", "WORD?"},
			want:  true,
		},
		{
			name:  "morelower",
			words: []string{"To", "HELL", "With", "georgia!"},
			want:  true,
		},
		{
			name:  "capsmoji",
			words: []string{`H0\/\/`, "'80U7", `T|-|3/\/\`, `d/-\\/\/6S?`},
			want:  true,
		},
		{
			name:  "lowermoji",
			words: []string{"p155", "on", "'3m!"},
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AllCapsDifferential(tt.words); got != tt.want {
				t.Errorf("AllCapsDifferential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCleanExtraPunc(t *testing.T) {
	tests := []struct {
		name string
		text string
		want string
	}{
		{
			name: "1",
			text: `"Yo…..what the fuck.." THEN IT CONTINUED and I was like "Yoooooooooooooooooooooooo……."`,
			want: `"Yo…what the fuck." THEN IT CONTINUED and I was like "Yoooooooooooooooooooooooo…"`,
		},
		{
			name: "2",
			text: "..THOSE NIGGAS…YOOOOOOO…….THOSE NIGGAS….",
			want: ".THOSE NIGGAS…YOOOOOOO…THOSE NIGGAS…",
		},
		{
			name: "3",
			text: "!!!Man ",
			want: "!Man",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CleanExtraPunc(tt.text); got != tt.want {
				t.Errorf("CleanExtraPunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
