package singleton

import "testing"

func TestSingleton(t *testing.T) {
	inst1 := GetInstance()
	inst1.Inc()
	inst1.Print()
	inst2 := GetInstance()
	inst2.Inc()
	inst2.Print()

}
