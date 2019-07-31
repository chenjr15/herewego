package factory

import "testing"

func TestFactory(t *testing.T) {
	speakable := GetSpeakable("Bird")
	speakable.Speak()
	speakable = GetSpeakable("Dog")
	speakable.Speak()

}
