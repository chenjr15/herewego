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

// BatchReadChannel 批量读取channel结果，每次最多读取n个，至少等待atLeastWait时间
// ok 表示是否成功读取到至少一个结果（即使后续channel关闭）
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
	// step 2. 判断是否需要继续读取
	remainingTime := atLeastWait - time.Since(start)
	if n < 2 && remainingTime <= time.Millisecond {
		// 时间已过，直接返回
		return results, true
	}
	ddlTimer := time.NewTimer(remainingTime)
	defer ddlTimer.Stop()
	// step 2. 批量读取结果 最多n-1个，超时后退出
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
