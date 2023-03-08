package main

import "fmt"

func changeSlice(s []int) {
	fmt.Println(s)
	s = s[:2]
	fmt.Println(s)
	// 调用后外层的s还是原来的长度，只有里面的长度会改变。说明这个slice类似于一个结构体，里面包含有一个数组指针
}

func main() {
	var s []int

	fmt.Printf("uninited slice var s []int \ns=%v, s==nil : %v\n", s, s == nil)
	emptySlice := []int{}
	fmt.Printf("uninited slice emptySlice := []int{} \nemptySlice=%v, emptySlice==nil : %v\n", emptySlice, emptySlice == nil)
	s = []int{0, 1, 2, 3}
	fmt.Println(s)
	changeSlice(s)
	fmt.Println(s)

}
