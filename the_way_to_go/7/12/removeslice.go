package main

import "fmt"

func main() {
	s := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Printf("sl:%v\n", s)
	s = RemoveSubSlice(s, 3, 5)
	fmt.Printf("sl:%v\n", s)
}
func RemoveSubSlice(s []byte, start, end int) (removed []byte) {
	if end <= start {
		return s
	}
	if start < 0 || end > len(s) {
		panic("Index out of boundary")
	}
	copy(s[start:], s[end:])
	removed = s[:len(s)-(end-start)]
	return
}
