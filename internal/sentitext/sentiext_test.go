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
				IsCapDiff: true,
			},
		},
		{
			name: "navy seal",
			text: `What the fuck did you just fucking say about me, you little bitch? I'll have you know I graduated top of my class in the Navy Seals, and I've been involved in numerous secret raids on Al-Quaeda, and I have over 300 confirmed kills. I am trained in gorilla warfare and I'm the top sniper in the entire US armed forces. You are nothing to me but just another target. I will wipe you the fuck out with precision the likes of which has never been seen before on this Earth, mark my fucking words. You think you can get away with saying that shit to me over the Internet? Think again, fucker. As we speak I am contacting my secret network of spies across the USA and your IP is being traced right now so you better prepare for the storm, maggot. The storm that wipes out the pathetic little thing you call your life. You're fucking dead, kid. I can be anywhere, anytime, and I can kill you in over seven hundred ways, and that's just with my bare hands. Not only am I extensively trained in unarmed combat, but I have access to the entire arsenal of the United States Marine Corps and I will use it to its full extent to wipe your miserable ass off the face of the continent, you little shit. If only you could have known what unholy retribution your little "clever" comment was about to bring down upon you, maybe you would have held your fucking tongue. But you couldn't, you didn't, and now you're paying the price, you goddamn idiot. I will shit fury all over you and you will drown in it. You're fucking dead, kiddo.`,
			wantS: &SentiText{
				IsCapDiff: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotS := Parse(tt.text)

			if !reflect.DeepEqual(gotS.IsCapDiff, tt.wantS.IsCapDiff) {
				t.Errorf("Parse() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}
