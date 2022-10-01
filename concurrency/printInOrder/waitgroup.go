package main

import (
	"sync"
)

var Println func(a ...interface{}) (n int, err error)

func setupPrint(f func(a ...interface{}) (n int, err error)) {
	Println = f
}

func main() {
	n := 4
	k := 102
	printNWaitGroup(n, k)
	printNMutex(n, k)
	printNChan(n, k)

}
func printNChan(n, k int) {
	cnt := 0
	chans := make([]chan struct{}, n)
	for i := range chans {
		chans[i] = make(chan struct{})
	}
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(idx int) {
			defer wg.Done()
			for range chans[idx] {
				if cnt < k {
					cnt++
					Println(idx, cnt)
					// 通知下一个协程
					chans[(idx+1)%n] <- struct{}{}
				} else {
					//关闭所有chan
					for i := 0; i < n; i++ {
						close(chans[i])
					}
				}
			}
		}(i)
	}
	chans[0] <- struct{}{}
	wg.Wait()

}
func printNMutex(n, k int) {
	// 基于互斥锁，保护一个临界变量
	// 然后每个线程抢到锁之后判断是否到自己打印的时机，是的话就加1并打印，不是的话就释放锁
	// 这个解法可能会导致抢锁很频繁，很多无效的判断
	mutex := sync.Mutex{}
	cnt := 0
	wg := sync.WaitGroup{}
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(idx int) {
			defer wg.Done()
			for {
				mutex.Lock()
				if cnt >= k {
					break
				}
				if cnt%n == idx {
					cnt++
					Println(idx, cnt)
				}
				mutex.Unlock()
			}
			// for break
			mutex.Unlock()

		}(i)
	}
	wg.Wait()
}

func printNWaitGroup(n, k int) {
	done := sync.WaitGroup{}
	done.Add(n)
	//	启动n个goroutine 依次打印1-k
	wgs := make([]*sync.WaitGroup, n)
	for i := 0; i < n; i++ {
		wgs[i] = &sync.WaitGroup{}
		// 所有wg加1
		wgs[i].Add(1)
	}
	wgs[0].Done()
	v := 0
	for i := 0; i < n; i++ {
		go func(i int) {
			defer done.Done()

			for {
				// Wait 当前
				wgs[i].Wait()
				if v < k {
					v++
					Println(i, v)
				}
				// Add 当前(下次才可以继续被Wait)
				wgs[i].Add(1)
				// Done 下一个
				wgs[(i+1)%n].Done()
				if v >= k {
					return
				}
			}

		}(i)
	}
	done.Wait()

}
