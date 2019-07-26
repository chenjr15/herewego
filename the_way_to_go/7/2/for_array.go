package main

import "fmt"

func main() {
	var arr [16]int
	for i := range arr {
		arr[i] = i
	}
	fmt.Println("array value :", arr)
}
