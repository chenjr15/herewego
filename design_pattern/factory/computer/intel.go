package computer

import "fmt"

type IntelFactory struct {
}
type IntelCPU struct{}

func (i IntelCPU) Calculate(things ...interface{}) {
	fmt.Println("Intel CPU, calculating:", things)
}

func (i IntelCPU) String() string {
	return "[Intel] CPU"
}

type IntelGPU struct{}

func (i IntelGPU) Display(things ...interface{}) {
	fmt.Println("Intel GPU, displaying:", things)
}

func (i IntelGPU) String() string {
	return "[Intel] GPU"
}

type IntelMemory struct{}

func (i IntelMemory) Storage(things ...interface{}) {
	fmt.Println("Intel Memory, storing:", things)
}

func (i IntelMemory) String() string {
	return "[Intel] Memory"
}

func (intel *IntelFactory) ProducingCPU() CPU {
	return &IntelCPU{}
}

func (intel *IntelFactory) ProducingGPU() GPU {
	return &IntelGPU{}
}

func (intel *IntelFactory) ProducingMemory() Memory {
	return &IntelMemory{}
}
