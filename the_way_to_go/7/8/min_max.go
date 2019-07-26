package main

import "fmt"

func main() {
	nums := []int{1, 2, 6, 23, 7, 0, -1}
	fmt.Println("min of", nums, "=", minSlice(nums))
	fmt.Println("max of", nums, "=", maxSlice(nums))
}
func minSlice(sl []int) (min int) {
	min = sl[0]
	for _, e := range sl {
		if e < min {
			min = e
		}
	}
	return
}
func maxSlice(sl []int) (max int) {
	max = sl[0]
	for _, e := range sl {
		if e > max {
			max = e
		}
	}
	return
}
