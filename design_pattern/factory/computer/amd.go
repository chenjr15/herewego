package computer

import "fmt"

type AMDFactory struct {
}
type AMDCPU struct{}

func (i AMDCPU) Calculate(things ...interface{}) {
	fmt.Println("AMD CPU, calculating:", things)
}

func (i AMDCPU) String() string {
	return "[AMD] CPU"
}

type AMDGPU struct{}

func (i AMDGPU) Display(things ...interface{}) {
	fmt.Println("AMD GPU, displaying:", things)
}

func (i AMDGPU) String() string {
	return "[AMD] GPU"
}

type AMDMemory struct{}

func (i AMDMemory) Storage(things ...interface{}) {
	fmt.Println("AMD Memory, storing:", things)
}

func (i AMDMemory) String() string {
	return "[AMD] Memory"
}

func (a *AMDFactory) ProducingCPU() CPU {
	return &AMDCPU{}
}

func (a *AMDFactory) ProducingGPU() GPU {
	return &AMDGPU{}
}

func (a *AMDFactory) ProducingMemory() Memory {
	return &AMDMemory{}
}
