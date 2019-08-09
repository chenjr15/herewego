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

// 将一棵树合并入另一颗二叉树，不改变其性质
func (bt *BST) MergeBST(anotherBT *BST) (father *Node) {
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

// InsertBST 将数据插入二叉排序树，不改变其性质
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
				bt = nil
			}

		} else {
			if bt.Left != nil {
				bt = bt.Left

			} else {
				bt.Left = node
				bt = nil
			}
		}
	}

	return node

}

// InsertBST 将数据插入二叉排序树，不改变其性质
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
				bt = nil
			}

		} else {
			if bt.Left != nil {
				bt = bt.Left

			} else {
				bt.Left = node
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

// SearchFather 根据指返回对应结点的父节点，查找失败返回nil
func (bt *BST) SearchFather(val int) (father *BST) {
	for bt != nil {
		if val == bt.Val {
			return father

		} else if val > bt.Val {
			father = bt
			bt = bt.Right

		} else if val < bt.Val {
			father = bt
			bt = bt.Left
		}
	}

	return nil

}

func (bt *BST) RemoveBST(val int) (toRemove *BST) {

	fmt.Printf("Remove %v\n", val)
	father := bt.SearchFather(val)
	if father == nil {
		return nil
	}
	fmt.Printf("father : %v, l: %v, r: %v \n", father, father.Left, father.Right)
	if father.Left != nil && father.Left.Val == val {
		toRemove = father.Left
		if toRemove.Left != nil {
			father.Left = toRemove.Left
		}

	} else {

		toRemove = father.Right
		fmt.Printf("ToRemove :%v ", toRemove)
		if toRemove.Left != nil {
			father.Right = toRemove.Left
		}
	}

	father.MergeBST(toRemove.Right)

	return toRemove

}
func (bt *BST) BSTString() string {
	result := ""
	appendStr := func(node *Node) {
		result += fmt.Sprintf("%v ", node)
	}
	bt.InOrderTraversal(appendStr)
	return result
}
