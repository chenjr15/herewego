package singleton

import (
	"sync"
	"sync/atomic"
)

// 懒汉式单例，第一次访问的时候才实例化对象，要注意线程安全
var lazyInst *single
var lazyLock sync.Mutex

func GetLazyInstanceMutex() *single {
	if lazyInst == nil {
		lazyLock.Lock()
		defer lazyLock.Unlock()
		if lazyInst == nil {
			lazyInst = new(single)
		}
	}
	return lazyInst
}

var lazyFlag uint32

func GetLazyInstanceAtomic() *single {
	// 用原子量做标记判断是否已经实例化
	if atomic.LoadUint32(&lazyFlag) == 1 {
		return lazyInst
	}

	lazyLock.Lock()
	defer lazyLock.Unlock()
	if lazyInst == nil {
		lazyInst = new(single)
		atomic.StoreUint32(&lazyFlag, 1)
	}

	return lazyInst
}

var once sync.Once

func GetLazyInstanceOnce() *single {
	// GetLazyInstanceAtomic 的标准库实现
	once.Do(func() {
		lazyInst = new(single)
	})

	return lazyInst
}
