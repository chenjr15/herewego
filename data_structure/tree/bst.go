package tree

// BST binary search tree
type BST = Node

// MakeBST makes a binary search tree, from the input values
func MakeBST(vals []int) *Node {

	bst := New(vals[0])
	for i := 1; i < len(vals); i++ {
		bst.InsertBST(vals[i])
	}

	return bst
}

// InsertBST 将数据插入二叉排序树，不改变其性质
func (bt *BST) InsertBST(val int) *Node {
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
