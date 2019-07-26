package main

import "fmt"

func main() {
	var hello = func() { fmt.Println("Hello world! Here is 匿名函数") }
	hello()
}
