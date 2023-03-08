package main

import "fmt"

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
