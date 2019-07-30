package method

import "testing"

// TestIncrease
func TestIncrease(t *testing.T) {

	var a Int
	t.Log(a)
	a.Increase(100)
	if a == 100 {
		t.Log(a)

	} else {
		t.Error("Failed, doesn't increased")
	}
}
