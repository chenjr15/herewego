package slice_test

import "testing"

func TestSumup(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}

	t.Log("sum of ", nums, "=", Sumup(nums...))
}

func Sumup(nums ...int) (sum int) {
	for _, item := range nums {
		sum += item
	}
	return

}
