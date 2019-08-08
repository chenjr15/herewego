package tree

import (
	"fmt"
	"testing"
)

func TestInOrder(t *testing.T) {
	tree := &Node{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	tree.AddLeft(1).AddRight(2)
	tree.InOrderTraversal(func(node *Node) { fmt.Println(node.Val) })

}
func TestLevelOrder(t *testing.T) {
	tree := &Node{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	tree.AddLeft(1).AddRight(2)
	tree.LevelOrderTraversal(func(node *Node, level int) { t.Log(level, node.Val) })

}
func TestTreePrint(t *testing.T) {
	tree := &Node{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	tree.AddLeft(1).AddRight(5)
	tree.AddRight(2).AddLeft(6)
	tree.AddLeft(3)
	tree.AddRight(4)
	s := tree.TreeString()
	fmt.Print(s)

}
