package singleton

import (
	"fmt"
	"sync"
)

type single struct {
	cnt int
	mu  sync.Mutex
}

// 饿汉式单例，无论是否使用，先提供好单例实例
var instance *single = new(single)

func GetInstance() *single {
	return instance
}

// GetInstanceUnsafe 返回单例的实例. 线程不安全
func GetInstanceUnsafe() *single {
	if instance == nil {
		instance = new(single)
	}
	return instance
}
func (inst *single) Print() {
	fmt.Printf("%p, %v\n", inst, inst.cnt)
}

func (inst *single) Inc() int {
	inst.mu.Lock()
	inst.cnt++
	inst.mu.Unlock()
	return inst.cnt
}
