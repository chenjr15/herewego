package slice_test

import "testing"

func TestSliceLen(t *testing.T) {
	s := []int{1, 2, 3}
	t.Log("len of s[n:n] = ", len(s[1:1]))
	t.Log("len of s[n:n+1] = ", len(s[1:1+1]))
}
