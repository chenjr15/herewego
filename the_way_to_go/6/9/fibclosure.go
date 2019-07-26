package main

import "fmt"

func main() {
	var f = fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(i, f())
	}
}
func fibonacci() func() int {

	var (
		count = 0
		last  [2]int
	)
	last[1] = 1
	return func() int {
		count++
		if count < 2 {

			return last[1]
		}
		last[0], last[1] = last[1], last[0]+last[1]
		count++
		return last[1]

	}

}
