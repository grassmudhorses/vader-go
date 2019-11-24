package sentitext

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		text      string
		wantcap   bool
		wantwords string
	}{
		{
			name:      "navy seal",
			text:      `What the fuck did you just fucking say about me, you little bitch? I'll have you know I graduated top of my class in the Navy Seals, and I've been involved in numerous secret raids on Al-Quaeda, and I have over 300 confirmed kills. I am trained in gorilla warfare and I'm the top sniper in the entire US armed forces. You are nothing to me but just another target. I will wipe you the fuck out with precision the likes of which has never been seen before on this Earth, mark my fucking words. You think you can get away with saying that shit to me over the Internet? Think again, fucker. As we speak I am contacting my secret network of spies across the USA and your IP is being traced right now so you better prepare for the storm, maggot. The storm that wipes out the pathetic little thing you call your life. You're fucking dead, kid. I can be anywhere, anytime, and I can kill you in over seven hundred ways, and that's just with my bare hands. Not only am I extensively trained in unarmed combat, but I have access to the entire arsenal of the United States Marine Corps and I will use it to its full extent to wipe your miserable ass off the face of the continent, you little shit. If only you could have known what unholy retribution your little "clever" comment was about to bring down upon you, maybe you would have held your fucking tongue. But you couldn't, you didn't, and now you're paying the price, you goddamn idiot. I will shit fury all over you and you will drown in it. You're fucking dead, kiddo.`,
			wantcap:   false,
			wantwords: `["What","the","fuck","did","you","just","fucking","say","about","me","you","little","bitch","I'll","have","you","know","I","graduated","top","of","my","class","in","the","Navy","Seals","and","I've","been","involved","in","numerous","secret","raids","on","Al-Quaeda","and","I","have","over","300","confirmed","kills","I","am","trained","in","gorilla","warfare","and","I'm","the","top","sniper","in","the","entire","US","armed","forces","You","are","nothing","to","me","but","just","another","target","I","will","wipe","you","the","fuck","out","with","precision","the","likes","of","which","has","never","been","seen","before","on","this","Earth","mark","my","fucking","words","You","think","you","can","get","away","with","saying","that","shit","to","me","over","the","Internet","Think","again","fucker","As","we","speak","I","am","contacting","my","secret","network","of","spies","across","the","USA","and","your","IP","is","being","traced","right","now","so","you","better","prepare","for","the","storm","maggot","The","storm","that","wipes","out","the","pathetic","little","thing","you","call","your","life","You're","fucking","dead","kid","I","can","be","anywhere","anytime","and","I","can","kill","you","in","over","seven","hundred","ways","and","that's","just","with","my","bare","hands","Not","only","am","I","extensively","trained","in","unarmed","combat","but","I","have","access","to","the","entire","arsenal","of","the","United","States","Marine","Corps","and","I","will","use","it","to","its","full","extent","to","wipe","your","miserable","ass","off","the","face","of","the","continent","you","little","shit","If","only","you","could","have","known","what","unholy","retribution","your","little","clever","comment","was","about","to","bring","down","upon","you","maybe","you","would","have","held","your","fucking","tongue","But","you","couldn't","you","didn't","and","now","you're","paying","the","price","you","goddamn","idiot","I","will","shit","fury","all","over","you","and","you","will","drown","in","it","You're","fucking","dead","kiddo"]`,
		},

		{
			name:      "navy seal but loud",
			text:      `WHAT THE FUCK DID YOU JUST FUCKING SAY ABOUT ME, YOU LITTLE BITCH? I'LL HAVE YOU KNOW I GRADUATED TOP OF MY CLASS IN THE NAVY SEALS, and I've been involved in numerous secret raids on Al-Quaeda, and I have over 300 confirmed kills. I am trained in gorilla warfare and I'm the top sniper in the entire US armed forces. You are nothing to me but just another target. I will wipe you the fuck out with precision the likes of which has never been seen before on this Earth, mark my fucking words. You think you can get away with saying that shit to me over the Internet? Think again, fucker. As we speak I am contacting my secret network of spies across the USA and your IP is being traced right now so you better prepare for the storm, maggot. The storm that wipes out the pathetic little thing you call your life. You're fucking dead, kid. I can be anywhere, anytime, and I can kill you in over seven hundred ways, and that's just with my bare hands. Not only am I extensively trained in unarmed combat, but I have access to the entire arsenal of the United States Marine Corps and I will use it to its full extent to wipe your miserable ass off the face of the continent, you little shit. If only you could have known what unholy retribution your little "clever" comment was about to bring down upon you, maybe you would have held your fucking tongue. But you couldn't, you didn't, and now you're paying the price, you goddamn idiot. I will shit fury all over you and you will drown in it. You're fucking dead, kiddo.`,
			wantcap:   true,
			wantwords: `["WHAT","THE","FUCK","DID","YOU","JUST","FUCKING","SAY","ABOUT","ME","YOU","LITTLE","BITCH","I'LL","HAVE","YOU","KNOW","I","GRADUATED","TOP","OF","MY","CLASS","IN","THE","NAVY","SEALS","and","I've","been","involved","in","numerous","secret","raids","on","Al-Quaeda","and","I","have","over","300","confirmed","kills","I","am","trained","in","gorilla","warfare","and","I'm","the","top","sniper","in","the","entire","US","armed","forces","You","are","nothing","to","me","but","just","another","target","I","will","wipe","you","the","fuck","out","with","precision","the","likes","of","which","has","never","been","seen","before","on","this","Earth","mark","my","fucking","words","You","think","you","can","get","away","with","saying","that","shit","to","me","over","the","Internet","Think","again","fucker","As","we","speak","I","am","contacting","my","secret","network","of","spies","across","the","USA","and","your","IP","is","being","traced","right","now","so","you","better","prepare","for","the","storm","maggot","The","storm","that","wipes","out","the","pathetic","little","thing","you","call","your","life","You're","fucking","dead","kid","I","can","be","anywhere","anytime","and","I","can","kill","you","in","over","seven","hundred","ways","and","that's","just","with","my","bare","hands","Not","only","am","I","extensively","trained","in","unarmed","combat","but","I","have","access","to","the","entire","arsenal","of","the","United","States","Marine","Corps","and","I","will","use","it","to","its","full","extent","to","wipe","your","miserable","ass","off","the","face","of","the","continent","you","little","shit","If","only","you","could","have","known","what","unholy","retribution","your","little","clever","comment","was","about","to","bring","down","upon","you","maybe","you","would","have","held","your","fucking","tongue","But","you","couldn't","you","didn't","and","now","you're","paying","the","price","you","goddamn","idiot","I","will","shit","fury","all","over","you","and","you","will","drown","in","it","You're","fucking","dead","kiddo"]`,
		},

		{
			name:      "emoji-laden",
			text:      `Just me and my 💕daddy💕, hanging out I got pretty hungry🍆 so I started to pout 😞 He asked if I was down ⬇for something yummy 😍🍆 and I asked what and he said he'd give me his 💦cummies!💦 Yeah! Yeah!💕💦 I drink them!💦 I slurp them!💦 I swallow them whole💦 😍 It makes 💘daddy💘 😊happy😊 so it's my only goal... 💕💦😫Harder daddy! Harder daddy! 😫💦💕 1 cummy💦, 2 cummy💦💦, 3 cummy💦💦💦, 4💦💦💦💦 I'm 💘daddy's💘 👑princess 👑but I'm also a whore! 💟 He makes me feel squishy💗!He makes me feel good💜! 💘💘💘He makes me feel everything a little should!~ 💘💘💘 👑💦💘Wa-What!💘💦👑`,
			wantcap:   false,
			wantwords: `["Just","me","and","my","💕","daddy","💕","hanging","out","I","got","pretty","hungry","🍆","so","I","started","to","pout","😞","He","asked","if","I","was","down","⬇","for","something","yummy","😍🍆","and","I","asked","what","and","he","said","he'd","give","me","his","💦","cummies","💦","Yeah","Yeah","💕💦","I","drink","them","💦","I","slurp","them","💦","I","swallow","them","whole","💦","😍","It","makes","💘","daddy","💘","😊","happy","😊","so","it's","my","only","goal","💕💦😫","Harder","daddy","Harder","daddy","😫💦💕","1","cummy","💦","2","cummy","💦💦","3","cummy","💦💦💦","4","💦💦💦💦","I'm","💘","daddy's","💘","👑","princess","👑","but","I'm","also","a","whore","💟","He","makes","me","feel","squishy","💗","He","makes","me","feel","good","💜","💘💘💘","He","makes","me","feel","everything","a","little","should","💘💘💘","👑💦💘","Wa-What","💘💦👑"]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS := Parse(tt.text)
			wordsp := *gotS.WordsAndEmotes
			words := make([]string, len(wordsp))
			for i, v := range wordsp {
				words[i] = v.Word
			}

			gotwords, err := json.Marshal(words)
			if err != nil {
				t.Error(err)
			}
			if string(gotwords) != string(tt.wantwords) {
				t.Errorf("Parse() = %v, want %v", string(gotwords), string(tt.wantwords))
			}

			if gotS.IsCapDiff != tt.wantcap {
				t.Errorf("Parse() = %v, want %v", gotS.IsCapDiff, tt.wantcap)
			}
		})
	}
}

func TestAllCapsDifferential(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  bool
	}{

		{
			name:  "lesscaps",
			words: []string{"What's", "the", "good", "word?", "To", "HELL", "With", "georgia!", "here is some more characters"},
			want:  false,
		},
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
			sentis := make([]SentiWord, len(tt.words))
			for i := range tt.words {
				sentis[i] = SentiWord{Word: tt.words[i], IsCaps: tt.words[i] == strings.ToUpper(tt.words[i])}
			}
			if got := allCapsDifferential(&sentis); got != tt.want {
				t.Errorf("AllCapsDifferential() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_splitEmojis(t *testing.T) {
	tests := []struct {
		name  string
		token string
		want  []string
	}{
		{
			name:  "stacked",
			token: "👑💦💘Wa-What!💘💦👑",
			want:  []string{"👑💦💘", "Wa-What!", "💘💦👑"},
		},
		{
			name:  "surround",
			token: "💕daddy💕",
			want:  []string{"💕", "daddy", "💕"},
		},
		{
			name:  "multi",
			token: "💘💘💘He",
			want:  []string{"💘💘💘", "He"},
		},
		{
			name:  "pre",
			token: "yummy💘💘💘",
			want:  []string{"yummy", "💘💘💘"},
		},
		{
			name:  "embedded",
			token: "squishy💗!He",
			want:  []string{"squishy", "💗", "!He"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitEmojis(tt.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitEmojis() = %v, want %v", got, tt.want)
			}
		})
	}
}
