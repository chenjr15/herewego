package main

import "testing"

func TestAppend(t *testing.T) {
	arr := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8}
	sl1 := arr[1:4]
	sl2 := arr[:8]
	t.Log(Append(sl1, sl2))
	t.Log(arr)
}
func Append(slice, data []byte) []byte {
	// if (cap(slice)- len(slice)) < len(data){
	// 	slice
	// }
	return append(slice, data...)
}
