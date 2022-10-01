package computer

import "fmt"

type Computer struct {
	CPU
	GPU
	Memory
}

func (c Computer) String() string {
	return fmt.Sprintf("Computer with:\n %s \n %s\n %s\n", c.CPU, c.GPU, c.Memory)
}

type CPU interface {
	// Calculate  计算万物
	Calculate(...interface{})
	fmt.Stringer
}
type GPU interface {
	// Display 显示一切
	Display(...interface{})
	fmt.Stringer
}
type Memory interface {
	// Storage 存储宇宙
	Storage(...interface{})
	fmt.Stringer
}

type Factory interface {
	ProducingCPU() CPU
	ProducingGPU() GPU
	ProducingMemory() Memory
}
