package singleton

import "fmt"

type Single struct {
	cnt int
}

var instance *Single

// GetInstanceUnsafe() 返回单例的实例. 不保证线程安全
func GetInstanceUnsafe() *Single {
	if instance == nil {
		instance = new(Single)
	}
	return instance
}
func (inst *Single) Print() {
	fmt.Printf("%p, %v\n", inst, inst.cnt)
}

func (inst *Single) Inc() *Single {
	inst.cnt += 1
	return inst
}
