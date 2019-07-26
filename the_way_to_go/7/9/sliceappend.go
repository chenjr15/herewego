package main

import "fmt"

func main() {
	sl := make([]int, 2, 10)
	fmt.Printf("len of %p is %d\n", sl, len(sl))
	sl = Enlarge(sl, 4)
	fmt.Printf("len of %p is %d\n", sl, len(sl))
	sl = Enlarge(sl, 2)
	fmt.Printf("len of %p is %d\n", sl, len(sl))

}
func Enlarge(s []int, factor int) (newS []int) {
	newLen := len(s) * factor
	if newLen > cap(s) {
		// here need to enlarge
		newS = make([]int, newLen)
		copy(newS, s)
	} else {
		newS = s
	}
	newS = newS[:newLen]
	return

}
