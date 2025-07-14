package main

import "fmt"
import "time"

func main() {
	var closedchan chan int
	closedchan = make(chan int, 1)

	closedchan <- 1
	close(closedchan)
	fmt.Println("recv from closed chan ", <-closedchan)
	// panic: send on closed channel
	// closedchan <- 1
	var nilchan chan int
	// 下面两句都会 fatal error: all goroutines are asleep - deadlock!
	nilchan <- 1
	<-nilchan
}

// BatchReadChannel 批量读取channel结果，函数第一个结果不受超时时间控制，剩下n-1个结果受超时时间控制。
// ok 为true表示是否成功读取到至少一个结果（即使后续channel关闭），ok 为false 则表示channel已关闭，且结果为空。
func BatchReadChannel[T any](channel <-chan T, n int, atLeastWait time.Duration) (results []T, ok bool) {
	if channel == nil {
		return nil, false
	}
	start := time.Now()
	// step 1. 至少读取一个结果
	result, ok := <-channel
	if !ok {
		return nil, false
	}
	results = append(results, result)
	// step 2.1 判断是否需要继续读取
	remainingTime := atLeastWait - time.Since(start)
	if n < 2 || remainingTime <= time.Millisecond {
		// 数量满足，或者时间已过，直接返回
		return results, true
	}
	ddlTimer := time.NewTimer(remainingTime)
	defer ddlTimer.Stop()
	// step 2.2 批量读取结果 最多n-1个，超时后退出
	for i := 1; i < n; i++ {
		select {
		case result, ok = <-channel:
			if !ok {
				// channel 已关闭，直接返回
				return results, true
			}
			results = append(results, result)
		case <-ddlTimer.C:
			// 超时退出
			return results, true
		}
	}
	return results, true
}
