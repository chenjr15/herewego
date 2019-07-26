package main

import "testing"

func TestSplit(t *testing.T) {
	var buf [10]byte
	for i := 0; i < len(buf); i++ {
		buf[i] = byte(i)
	}
	t.Log(Split(buf[:], 5))
}
func Split(buf []byte, n int) (sl1, sl2 []byte) {

	return buf[:n], buf[n:]
}
