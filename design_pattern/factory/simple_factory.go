package factory

import "fmt"

type Speakable interface {
	Speak()
}

type Dog struct {
}

func (dog *Dog) Speak() {
	fmt.Println("Wang!")
}

type Bird struct {
}

func (dog *Bird) Speak() {
	fmt.Println("Jiujiu!")
}
func GetSpeakable(name string) (speakable Speakable) {
	switch name {
	case "Bird":
		speakable = new(Bird)
	case "Dog":
		speakable = new(Dog)
	default:
		speakable = nil
		panic("Unknown type")

	}
	return
}
