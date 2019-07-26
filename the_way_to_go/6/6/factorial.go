package main

import "fmt"

func main() {
	for i := uint64(0); i <= 30; i++ {
		fmt.Printf("%v! = %v\n", i, factorial(i))
	}
}

func factorial(n uint64) (result uint64) {
	if n == 0 {
		return 1
	}
	return factorial(n-1) * n
}
