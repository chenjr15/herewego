package strategy

import "testing"

func TestOperation(t *testing.T) {
	operators := []Operator{&Add{}, &Minus{}, &Multi{}, &Div{}}
	operation := Operation{op: nil}

	a, b := 6, 3
	for _, op := range operators {
		operation.op = op
		r := operation.Operate(a, b)
		t.Logf("%d %s %d = %d  ", a, op.GetName(), b, r)

	}

}
