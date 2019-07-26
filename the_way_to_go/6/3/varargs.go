package main

import "fmt"

func main() {
	lines := []string{"first", "second", "third"}
	// var argument can only be slice , not an array
	NewLinePrint(lines...)
}
func NewLinePrint(contents ...string) {
	for _, line := range contents {
		fmt.Println(line)
	}
}
