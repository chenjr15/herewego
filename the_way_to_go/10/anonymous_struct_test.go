package method

import "testing"

type A struct {
	f float32
	int
	string
}

func TestAnonumousStruct(t *testing.T) {
	a := A{0.2, 3, "Struct"}
	t.Log(a)

}
