package main

import "fmt"

func main() {
	sl1 := []byte{1, 2, 3}
	sl2 := []byte{7, 8, 9}
	printSlice(sl1)
	printSlice(sl2)
	printSlice(InsertSlice(sl1, sl2, 1))
	printSlice(sl1)
	printSlice(sl2)
	fmt.Println("vim-go")
}
func printSlice(x []byte) {
	fmt.Printf("len=%d  cap=%d   slice=%v\n", len(x), cap(x), x)

}

/// InsertSlice: insert a slice to another slice.
func InsertSlice(dest, to_insert []byte, pos int) (inserted []byte) {
	front, behind := dest[:pos], dest[pos:]
	newlen := (len(dest) + len(to_insert))

	if cap(dest) < newlen {
		// need to enlarge
		inserted = make([]byte, newlen)
		copy(inserted, front)
	} else {
		// don't need to enlarge, just use the old one,
		// need to relice it to capable of new size
		inserted = dest[:newlen]
	}
	// copy the second part of origin slice
	copy(inserted[newlen-len(behind):], behind)
	copy(inserted[pos:], to_insert)

	return
}
