package tree

import (
	"fmt"
	"testing"
)

var tree *Node

func Print(node *Node) {
	fmt.Println(node.Val)
}
func init() {
	tree = New(0)
	tree.AddLeft(1).AddRight(4)
	tree.AddRight(2).AddLeft(5)
	tree.AddLeft(3)
	tree.AddRight(6)

}

func TestInOrder(t *testing.T) {
	tree := New(0)
	tree.AddLeft(1).AddRight(2)
	tree.InOrderTraversal(Print)

}
func TestLevelOrder(t *testing.T) {
	tree := New(0)
	tree.AddLeft(1).AddRight(2)
	tree.LevelOrderTraversal(func(node *Node, level int) { t.Log(level, node.Val) })

}
func TestTreePrint(t *testing.T) {
	tree.AttachLeft(New(9))
	s := tree.TreeString()
	fmt.Print(s)

}
func TestPreOrder(t *testing.T) {

	tree.PreOrderTraversal(Print)

}
func TestPostOrder(t *testing.T) {

	tree.PostOrderTraversal(Print)

}
