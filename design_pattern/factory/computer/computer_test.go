package computer

import (
	"fmt"
	"testing"
)

func TestAllIntel(t *testing.T) {
	var computer Computer
	var factory Factory
	factory = &IntelFactory{}
	computer.CPU = factory.ProducingCPU()
	computer.GPU = factory.ProducingGPU()
	computer.Memory = factory.ProducingMemory()
	fmt.Println(computer)
}
