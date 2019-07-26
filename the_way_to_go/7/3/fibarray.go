package main

import "fmt"

var memo [51]uint

func main() {

	for i := 0; i < 50; i++ {

		fmt.Println(i, fib(i))
	}
}
func fib(n int) uint {
	if memo[n] != 0 {
		return memo[n]
	}
	if n <= 2 {
		return uint(1)
	}
	memo[n] = fib(n-1) + fib(n-2)
	return memo[n]
}
