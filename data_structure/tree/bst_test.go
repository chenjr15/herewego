package tree

import (
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	n := 10
	vals := make([]int, n)
	flag := -1
	for i := range vals {
		flag = -1 * flag
		vals[i] = i + 3*flag + 3
	}
	bst := MakeBST(vals)
	t.Log(vals)
	result := ""
	visit := func(node *Node) {

		result += fmt.Sprintf("%v", node)
	}
	bst.InOrderTraversal(visit)
	t.Log(result)
}
