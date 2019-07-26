package slice_test

import "testing"

func TestDouble(t *testing.T) {
	items := [...]int{10, 20, 30, 40, 50}
	for _, item := range items {
		item *= 2
	}
	t.Log(items)
}
func TestDoubleWorks(t *testing.T) {
	items := [...]int{10, 20, 30, 40, 50}
	for ix, _ := range items {
		items[ix] *= 2
	}
	t.Log(items)
}
