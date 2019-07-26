package slice_test

import "testing"

// slice 引用原数组，append在原数组上操作
func TestSliceModify(t *testing.T) {
	var arr [10]int
	t.Log("arr len:", len(arr), "cap:", cap(arr), "\nvalue:", arr)
	sle := arr[4:6]
	t.Log("slice len:", len(sle), "cap:", cap(sle))
	s2 := append(sle, 2)
	t.Log("after func call value\n", arr)
	t.Log("s2 len:", len(s2), "cap:", cap(s2))
	t.Log("after func call new slice\n", s2)
	s3 := sle[:5]
	s3[len(s3)-1] = 3
	t.Log("after func call value\n", arr)
	t.Log("s3 len:", len(s3), "cap:", cap(s3))
	t.Log("after func call new slice\n", s3)
}
func TestSliceMake(t *testing.T) {
	s := make([]byte, 5)
	t.Log("slice len:", len(s), "cap:", cap(s))
	s = s[2:4]
	t.Log("slice len:", len(s), "cap:", cap(s))
}
