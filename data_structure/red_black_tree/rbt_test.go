package rbtree

import "testing"

func TestAAAA(t *testing.T) {
	tree := New(6)
	tree.Color = BLACK
	node := tree.AddLeft(1)
	node.Color = RED
	node = node.AddRight(3)

	node.Color = BLACK
	node = node.AddRight(5)
	node.Color = BLACK

	node = tree.AddRight(8)
	node.Color = RED
	node.AddLeft(7).Color = BLACK
	node = tree.AddRight(10)
	node.Color = BLACK
	node.AddLeft(9).Color = BLACK
	node = tree.AddRight(12)
	node.Color = RED
	node = tree.AddRight(14)
	node.Color = BLACK

	t.Logf("%v", tree)
	print := func() {

		result := "\n"
		currentLevel := 0
		visit := func(node *Node, level int) {
			if level != currentLevel {
				result += "\n"
			}
			currentLevel = level

			result += node.String() + "^" + node.Parent.String() + " "
		}
		tree.LevelOrderTraversal(visit)
		t.Log(result)
	}
	print()

	node = tree.Search(8)
	tree.LeftRotate(node)
	print()

}
