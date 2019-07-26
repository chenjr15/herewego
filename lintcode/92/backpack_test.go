package backpack

import "testing"

func TestBackpack(t *testing.T) {
	test := func(testcase []int, backSize, correct int) {

		a := Backpack(backSize, testcase)
		t.Log(a)
		if a != correct {
			t.Error("failed with", testcase, "correct is", 9, "got", a)
		}

	}
	testdata := []int{3, 4, 8, 5}
	test(testdata, 10, 9)
	//backpack.Backpack(5, make([]int, 5))

}
