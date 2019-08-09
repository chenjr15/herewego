package rbtree

// InsertRBT 添加节点到红黑树中
func (rbt RBT) InsertRBT(node *Node) *Node {
	rbt.InsertBST(node)
	node.Color = RED
	rbt.FixUP(node)

	return node
}

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
	if node == nil || node.Left == nil {
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
	if y.Right != nil {

		y.Right.Parent = node
	}
	y.Right = node
	node.Parent = y
	if node == rbt {
		return y
	}
	return rbt
}

// FixUP 修复红黑树的关系
func (rbt RBT) FixUP(node *Node) *Node {
	if rbt == nil {
		return rbt
	}
	// 确保该结点当前树上的结点
	node = rbt.Search(node.Key)
	// 确保结点有父节点，因为它会取代当前节点
	if node == nil || node.Parent == nil {
		return rbt
	}

	for node.Parent!=nil && node.Parent.Color == RED {
		if node.Parent == node.Parent.Parent.Left {
			y := node.Parent.Parent.Left
			if y.Color == RED {
				node.Parent.Color = BLACK
				y.Color = BLACK
				node.Parent.Parent.Color = RED
				node = node.Parent.Parent
			} else if node == node.Parent.Right {
				node = node.Parent
				rbt.LeftRotate(node)
			}
			if node.Parent != nil {
				node.Parent.Color = BLACK
				if node.Parent.Parent == nil{
					break 
				}
				node.Parent.Parent.Color = RED
				rbt.RightRotate(node.Parent.Parent)

			}
		} else {
			y := node.Parent.Parent.Right
			if y.Color == RED {
				node.Parent.Color = BLACK
				y.Color = BLACK
				node.Parent.Parent.Color = RED
				node = node.Parent.Parent
			} else if node == node.Parent.Left {
				node = node.Parent
				rbt.RightRotate(node)
			}
			if node.Parent != nil {
				node.Parent.Color = BLACK
				if node.Parent.Parent == nil{
					break 
				}
				node.Parent.Parent.Color = RED
				rbt.LeftRotate(node.Parent.Parent)
			}
		}

	}
	rbt.Color = BLACK
	return rbt

}
