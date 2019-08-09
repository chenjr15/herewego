package rbtree

// LeftRotate 左旋，left down ， 不会管颜色
func (rbt RBT) LeftRotate(node *Node) *Node {
	if rbt == nil {
		return rbt
	}
	// 确保该结点当前树上的结点
	node = rbt.Search(node.Key)
	// 确保结点有右节点，因为它会取代当前节点
	if node == nil || node.Right == nil {
		return rbt
	}

	y := node.Right
	p := node.Parent
	if p.Left == node {
		p.Left = y
	} else {
		p.Right = y
	}
	y.Parent = p
	node.Right = y.Left
	if y.Left != nil {

		node.Right.Parent = node
	}
	y.Left = node
	node.Parent = y
	if node == rbt {
		return y
	}
	return rbt

}

// RightRotate 右旋，right down，不会管颜色
func (rbt RBT) RightRotate(node *Node) *Node {
	if rbt == nil {
		return rbt
	}
	// 确保该结点当前树上的结点
	node = rbt.Search(node.Key)
	// 确保结点有左节点，因为它会取代当前节点
	if node == nil || node.Left== nil {
		return rbt
	}

	y := node.Left
	p := node.Parent
	if p.Left == node {
		p.Left = y
	} else {
		p.Right = y
	}
	y.Parent = p
	node.Left = y.Right
	if y.Right!= nil {

		y.Right.Parent = node
	}
	y.Right= node
	node.Parent = y
	if node == rbt {
		return y
	}
	return rbt
}

// FixUP 修复红黑树的关系
func (rbt RBT) FixUP(node *Node) {

}
