package tree

import "fmt"

// BST binary search tree
type BST = Node

// MakeBST makes a binary search tree, from the input values
func MakeBST(vals []int) *Node {

	bst := New(vals[0])
	for i := 1; i < len(vals); i++ {
		bst.AddBST(vals[i])
	}

	return bst
}

// MergeBST 将一棵树合并入另一颗二叉树，不改变其性质
func (bt *BST) MergeBST(anotherBT *BST) (parent *Node) {
	if bt == nil {
		return anotherBT
	}
	add := func(n *Node, level int) {
		n.Left = nil
		n.Right = nil
		bt.InsertBST(n)

	}
	anotherBT.LevelOrderTraversal(add)

	return bt

}

// InsertBST 将结点插入二叉排序树，不改变其性质
func (bt *BST) InsertBST(node *Node) *Node {
	if bt == nil {
		return nil
	}
	if bt == node {
		return bt
	}
	for bt != nil {
		if node.Val > bt.Val {
			if bt.Right != nil {
				bt = bt.Right
			} else {
				bt.Right = node
				node.Parent = bt
				bt = nil
			}

		} else {
			if bt.Left != nil {
				bt = bt.Left

			} else {
				bt.Left = node
				node.Parent = bt
				bt = nil
			}
		}
	}

	return node

}

// AddBST 将数据插入二叉排序树，不改变其性质
func (bt *BST) AddBST(val int) *Node {
	if bt == nil {
		return nil
	}
	node := New(val)
	for bt != nil {
		if val > bt.Val {
			if bt.Right != nil {
				bt = bt.Right
			} else {
				bt.Right = node
				node.Parent = bt
				bt = nil
			}

		} else {
			if bt.Left != nil {
				bt = bt.Left

			} else {
				bt.Left = node
				node.Parent = bt
				bt = nil
			}
		}
	}

	return node

}

// Search 根据指返回对应结点，查找失败返回nil
func (bt *BST) Search(val int) (node *BST) {
	for bt != nil {
		if val == bt.Val {
			node = bt
			break
		} else if val > bt.Val {
			bt = bt.Right

		} else if val < bt.Val {
			bt = bt.Left
		}
	}

	return

}

// RemoveBST 删除结点. 无法删除根节点
func (bt *BST) RemoveBST(val int) (toRemove *BST) {

	toRemove = bt.Search(val)
	if toRemove == nil {
		return nil
	}
	parent := toRemove.Parent

	if parent == nil {
		return toRemove
	}
	if parent.Left != nil && parent.Left == toRemove {
		if toRemove.Left != nil {
			parent.Left = toRemove.Left
			toRemove.Left.Parent = parent
		} else {
			parent.Left = nil
		}

	} else {

		if toRemove.Left != nil {
			parent.Right = toRemove.Left
			toRemove.Left.Parent = parent
		} else {
			parent.Right = nil
		}
	}

	parent.MergeBST(toRemove.Right)

	return toRemove

}

// BSTString 按层输出
func (bt *BST) BSTString() string {
	result := ""
	appendStr := func(node *Node) {
		result += fmt.Sprintf("%v ", node)
	}
	bt.InOrderTraversal(appendStr)
	return result
}
