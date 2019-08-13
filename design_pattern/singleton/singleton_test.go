package singleton

import "testing"

func TestSingleton(t *testing.T) {
	inst1 := GetInstanceUnsafe()
	inst1.Inc()
	inst1.Print()
	inst2 := GetInstanceUnsafe()
	inst2.Print()

}
