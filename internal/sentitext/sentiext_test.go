package sentitext

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name  string
		text  string
		wantS *SentiText
	}{
		{
			name: "the ultimate expression of emotion copypasta",
			text: `!!!Man this shit is so wrong in so many motherfucking levels yo…I was talking to one of my white friends and he sent me 3 videos with the name only labeled "Boku" I said to this dude, What's this shit? He just giggled and said "Just watch them and MAKE SURE NOBODY IS AROUND YOU WHEN WATCHING IT!" Then I thought it was some weird porno or some strange shit but as I watched the first video, I was like "Yo…..what the fuck.." THEN IT CONTINUED and I was like "Yoooooooooooooooooooooooo……." THEN THEY GOT IN THE MOTHERFUCKING CAR AND THEN I SAID "YYYYYYYYYYYYYYYYYYYYYYYYYYYYYOOOOOOOOOO OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO!!!!!" I couldn't fucking believe what I just saw, It was like Satan gave me his porno collection, shit was so disturbing..YET I COULDN'T STOP WATCHING IT, THEN VIDEO TWO AND IT WAS TWO OF THEM…..THOSE NIGGAS…YOOOOOOO…….THOSE NIGGAS….AND THAT GIRL SAW THEM THEN SHE…YYYYYYYYYOOOOOOOOOOOOOOOOOOOOOO… THEN THAT NIGGA TOOK THAT DOG TOY THEN YYYYYYYYYYYYYOOOOOOOOOOOOOOOOOOOOOOOOOOO OOOOOOOOOOOOOOOOOOOOOOOOOO……..IT WAS LIKE YOUR BITCH WANTED TO HAVE SEX WITH YOU BUT SHE WANTED TO SOMETHING "DIFFERENT" AND IT WAS SO FUCKED UP AND CREEPY, YOU JUST…KEPT WATCHING IT…AND THAT'S WHAT I FUCKING DID!!!!! THEN I SAW VIDEO THREE…THREE NIGGAS…THRRREEEEE!!!!!! IT…WAS…THHHHHHRRRRRRRRRREEEEEEEEEEEE EEEEEEEEEEEEEEEEEEEEEEE!!!!!!!!!! AND COCO WAS HIS NAME NIGGA, COCO WAS HIS MOTHERFUCKING NAME!!!!!! OH MY GOD,I AIN'T GOING TO HEAVEN NIGGAS, I ALREADY SOLD MY SOUL TO LUCIFER!`,
			wantS: &SentiText{
				Text:      `!!!Man this shit is so wrong in so many motherfucking levels yo…I was talking to one of my white friends and he sent me 3 videos with the name only labeled "Boku" I said to this dude, What's this shit? He just giggled and said "Just watch them and MAKE SURE NOBODY IS AROUND YOU WHEN WATCHING IT!" Then I thought it was some weird porno or some strange shit but as I watched the first video, I was like "Yo…..what the fuck.." THEN IT CONTINUED and I was like "Yoooooooooooooooooooooooo……." THEN THEY GOT IN THE MOTHERFUCKING CAR AND THEN I SAID "YYYYYYYYYYYYYYYYYYYYYYYYYYYYYOOOOOOOOOO OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO!!!!!" I couldn't fucking believe what I just saw, It was like Satan gave me his porno collection, shit was so disturbing..YET I COULDN'T STOP WATCHING IT, THEN VIDEO TWO AND IT WAS TWO OF THEM…..THOSE NIGGAS…YOOOOOOO…….THOSE NIGGAS….AND THAT GIRL SAW THEM THEN SHE…YYYYYYYYYOOOOOOOOOOOOOOOOOOOOOO… THEN THAT NIGGA TOOK THAT DOG TOY THEN YYYYYYYYYYYYYOOOOOOOOOOOOOOOOOOOOOOOOOOO OOOOOOOOOOOOOOOOOOOOOOOOOO……..IT WAS LIKE YOUR BITCH WANTED TO HAVE SEX WITH YOU BUT SHE WANTED TO SOMETHING "DIFFERENT" AND IT WAS SO FUCKED UP AND CREEPY, YOU JUST…KEPT WATCHING IT…AND THAT'S WHAT I FUCKING DID!!!!! THEN I SAW VIDEO THREE…THREE NIGGAS…THRRREEEEE!!!!!! IT…WAS…THHHHHHRRRRRRRRRREEEEEEEEEEEE EEEEEEEEEEEEEEEEEEEEEEE!!!!!!!!!! AND COCO WAS HIS NAME NIGGA, COCO WAS HIS MOTHERFUCKING NAME!!!!!! OH MY GOD,I AIN'T GOING TO HEAVEN NIGGAS, I ALREADY SOLD MY SOUL TO LUCIFER!`,
				IsCapDiff: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS := Parse(tt.text)
			if !reflect.DeepEqual(gotS.Text, tt.wantS.Text) {
				t.Errorf("Parse() = %v, want %v", gotS, tt.wantS)
			}

			if !reflect.DeepEqual(gotS.IsCapDiff, tt.wantS.IsCapDiff) {
				t.Errorf("Parse() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
