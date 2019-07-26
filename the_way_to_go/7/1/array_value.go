package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println("init value", arr)
	modifyArray(arr)
	fmt.Println("after func call value", arr)
}
func modifyArray(arr_inside [3]int) {
	arr_inside[0] = 999
	fmt.Println("inside func call value", arr_inside)
}
