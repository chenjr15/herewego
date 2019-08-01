package method

import "testing"

func TestGiveRaise(t *testing.T) {
	e := employee{100}
	t.Log(e)
	e.giveRaise(0.2)
	t.Log(e)
	if e.salary != 120.0 {
		t.Errorf("Wrong salary:%v", e)
	}

}
