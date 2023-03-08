package main

import "fmt"

func main() {

	ch := make(chan struct{})

	go func() {
		close(ch)
		ch <- struct{}{}
	}()

	i := 0
	for range ch {
		i++
	}
	fmt.Printf("%d", i)
}
