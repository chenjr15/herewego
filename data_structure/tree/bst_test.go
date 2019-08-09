package tree

import (
	"fmt"
	"testing"
)

var bst *BST

var vals []int
var fathers []int

func TestMerge(t *testing.T) {
	t1 := MakeBST([]int{5, 4, 10, 8, 11, 7, 9})
	t2 := MakeBST([]int{2, 1, 3})
	fmt.Println(t1.BSTString())
	fmt.Println(t2.BSTString())
	t1.MergeBST(t2)
	t.Log(t1.BSTString())
	t.Log("\n", t1.TreeString())

}
func TestRemove(t *testing.T) {
	t1 := MakeBST([]int{5, 4, 10, 8, 11, 7, 9})
	t.Log("\n", t1.BSTString())
	remove := []int{1, 4, 5, 10}
	for _, k := range remove {

		v := t1.RemoveBST(k)

		t.Log(v)
		t.Log("\n", t1.BSTString())
	}

}

func TestBST(t *testing.T) {
	t.Log(vals)
	result := ""
	visit := func(node *Node) {

		result += fmt.Sprintf("%v", node)
	}
	bst.InOrderTraversal(visit)
	t.Log(result)

}
func TestSearch(t *testing.T) {

	for _, val := range vals {
		node := bst.Search(val)
		if node == nil || node.Val != val {
			t.Errorf("%v misatch got :%v", val, node)
		}
	}
	notExists := []int{0, 2, 4, -2}
	for _, val := range notExists {
		node := bst.Search(val)
		if node != nil {
			t.Errorf("%v misatch got :%v", val, node)
		}
	}

	for i, val := range vals {
		node := bst.SearchFather(val)
		if node == nil || node.Val != fathers[i] {

			if node == nil && fathers[i] == -1 {
				break

			}
			t.Errorf("%v misatch got :%v", val, node)
		}
	}

}
func init() {
	n := 10
	vals = make([]int, n)
	flag := -1
	for i := range vals {
		flag = -1 * flag
		vals[i] = i + 3*flag + 3
	}
	bst = MakeBST(vals)
	fathers = []int{-1, 6, 6, 1, 8, 3, 10, 8, 12, 10}
}
