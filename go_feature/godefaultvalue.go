package main

import "fmt"

const TMP = "%s \t\t -> %v,%#v \t v == nil :%v \n"

func main() {
	var varSlice []int
	var varMap map[int]int

	fmt.Printf(TMP, "var []int \t ", varSlice, varSlice, varSlice == nil)
	fmt.Printf(TMP, "var map[int]int", varMap, varMap, varMap == nil)

	newSlice := new([]int)
	fmt.Printf(TMP, "new([]int)\t", newSlice, newSlice, newSlice == nil)
	newMap := new(map[int]int)
	fmt.Printf(TMP, "new(map[int]int)", newMap, newMap, newMap == nil)

}
