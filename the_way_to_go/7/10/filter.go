package main

import "fmt"

func main() {
	biggerThan10 := func(n int) bool {
		return n > 10
	}
	sl := []int{1, 2, 3, 4, 5, 11, 22, 23, 43, 5, 436, 54, 765, 8}

	fmt.Println("origin :", sl)
	fmt.Println(Filter(sl, biggerThan10))
}
func Filter(s []int, fn func(int) bool) (filtered []int) {

	for _, item := range s {
		if fn(item) {
			filtered = append(filtered, item)
		}
	}
	return
}
