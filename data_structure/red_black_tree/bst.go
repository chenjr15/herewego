package rbtree

import (
	"container/list"
	"fmt"
)

// Color color type of red black tree
type Color bool

// BLACK for color black of rbt
const BLACK = Color(true)

// RED for color red of rbt
const RED = Color(false)

// Node struct of tree
type Node struct {
	// Key use to compare
	Key    int
	Left   *Node
	Right  *Node
	Parent *Node
	Color
}

// RBT Red Black Tree
type RBT = *Node

// New Return a new Node with key
func New(key int) *Node {

	return &Node{
		Key:    key,
		Left:   nil,
		Right:  nil,
		Parent: nil,
		Color:  RED,
	}
}

// LevelOrderTraversal 层次遍历
// 用队列实现，先把头节点入队， 然后取队头，访问队头，再把对头的左右孩子入队, 循环至队列空
func (t *Node) LevelOrderTraversal(visit func(*Node, int)) {
	if t == nil {
		return
	}
	level := 0
	lastLevel := 0
	queue := list.New()
	if t != nil {
		queue.PushBack(t)
		lastLevel = 1
	}

	for lastLevel != 0 {
		thisLevel := lastLevel
		lastLevel = 0
		for thisLevel != 0 {
			t = (queue.Remove(queue.Front())).(*Node)
			thisLevel--
			if t.Left != nil {
				queue.PushBack(t.Left)
				lastLevel++
			}
			if t.Right != nil {
				queue.PushBack(t.Right)
				lastLevel++
			}
			visit(t, level)
		}
		level++

	}

}

// PreOrderTraversal 先序遍历
func (t *Node) PreOrderTraversal(visit func(*Node)) {
	if t == nil {
		return
	}
	stack := make([]*Node, 0)
	p := t
	for p != nil || len(stack) != 0 {
		if p != nil {
			visit(p)
			if p.Right != nil {
				stack = append(stack, p.Right)
			}
			if p.Left != nil {
				stack = append(stack, p.Left)
			}
			p = nil

		} else {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

	}
}

// PostOrderTraversal 后序遍历
// 双栈法， 稍微修改先序遍历，按照 根右左 的顺序将元素如打印栈，再不断弹出打印栈中的元素访问即可
func (t *Node) PostOrderTraversal(visit func(*Node)) {
	if t == nil {
		return
	}
	stack := make([]*Node, 0)
	result := make([]*Node, 0)
	p := t
	for p != nil || len(stack) != 0 {
		if p != nil {
			result = append(result, p)
			if p.Left != nil {
				stack = append(stack, p.Left)
			}
			if p.Right != nil {
				stack = append(stack, p.Right)
			}

			p = nil

		} else {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

	}
	for i := len(result) - 1; i > -1; i-- {
		visit(result[i])

	}
}

// InOrderTraversal 中序遍历
// 用栈实现，先碰到元素先入栈，然后往左边走，重复直到左边为空，然后出栈并访问，然后往右边走，重复前面的入栈往左边走
func (t *Node) InOrderTraversal(visit func(*Node)) {
	if t == nil {
		return
	}
	stack := make([]*Node, 0)
	p := t
	for p != nil || len(stack) != 0 {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if p == nil && len(stack) != 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			visit(p)
			p = p.Right
		}
	}
	return

}

// AttachLeft go to the end of left and attach a new node, then return the father of new node
func (t *Node) AttachLeft(node *Node) *Node {
	if t == nil {
		return nil
	}
	for t.Left != nil {
		t = t.Left
	}
	t.Left = node
	node.Parent = t

	return t

}

// AddLeft go to the end of left and add a new node, then return new node
func (t *Node) AddLeft(key int) *Node {
	if t == nil {
		return nil
	}
	for t.Left != nil {
		t = t.Left
	}
	t.Left = &Node{
		Key:    key,
		Left:   nil,
		Right:  nil,
		Parent: t,
	}
	return t.Left

}

// AttachRight go to the end of left and attach a new node, then return the father of new node
func (t *Node) AttachRight(node *Node) *Node {
	if t == nil {
		return nil
	}
	for t.Right != nil {
		t = t.Right
	}
	t.Right = node
	node.Parent = t

	return t

}

// AddRight go to the end of right and add a new node, then return new node
func (t *Node) AddRight(key int) *Node {
	if t == nil {
		return nil
	}
	for t.Right != nil {
		t = t.Right
	}
	t.Right = &Node{
		Key:    key,
		Left:   nil,
		Right:  nil,
		Parent: t,
	}
	return t.Right

}

// TreeString return preety format string
func (t *Node) TreeString() (result string) {
	if t == nil {
		return ""
	}
	leftPos := make(map[*Node]int)
	currentLevel := 0
	toLeft := 1
	for p := t.Left; p != nil; {
		p = p.Left
		toLeft++
	}
	leftPos[t] = toLeft
	levelOffset := 0

	visit := func(node *Node, level int) {

		if level != currentLevel {
			result += "\n"
			levelOffset = 0
		}
		currentLevel = level
		if levelOffset != 0 {
			levelOffset++
		}
		levelOffset = leftPos[node] - levelOffset

		margin := []byte{}
		for i := 0; i < 3*levelOffset; i++ {
			margin = append(margin, ' ')
		}

		result += fmt.Sprintf("%s%v", string(margin), node)
		if node == nil {
			return
		}
		if node.Left != nil {
			leftPos[node.Left] = leftPos[node] - 1
		}
		if node.Right != nil {
			leftPos[node.Right] = leftPos[node] + 1

		}

	}
	t.LevelOrderTraversal(visit)
	result += "\n"

	return result
}

// MakeBST makes a binary search tree, from the input keyues
func MakeBST(keys []int) *Node {

	bst := New(keys[0])
	for i := 1; i < len(keys); i++ {
		bst.AddBST(keys[i])
	}

	return bst
}

// MergeBST 将一棵树合并入另一颗二叉树，不改变其性质
func (bt *Node) MergeBST(anotherBT *Node) (parent *Node) {
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
func (bt *Node) InsertBST(node *Node) *Node {
	if bt == nil {
		return nil
	}
	if bt == node {
		return bt
	}
	for bt != nil {
		if node.Key > bt.Key {
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
func (bt *Node) AddBST(key int) *Node {
	if bt == nil {
		return nil
	}
	node := New(key)
	for bt != nil {
		if key > bt.Key {
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
func (bt *Node) Search(key int) (node *Node) {
	for bt != nil {
		if key == bt.Key {
			node = bt
			break
		} else if key > bt.Key {
			bt = bt.Right

		} else if key < bt.Key {
			bt = bt.Left
		}
	}

	return

}

// RemoveBST 删除结点. 无法删除根节点
func (bt *Node) RemoveBST(key int) (toRemove *Node) {

	toRemove = bt.Search(key)
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
func (bt *Node) BSTString() string {
	result := ""
	appendStr := func(node *Node) {
		result += fmt.Sprintf("%v ", node)
	}
	bt.InOrderTraversal(appendStr)
	return result
}
func (rbt *Node) String() (s string) {
	if rbt == nil {
		return "[-]*"
	}
	s = fmt.Sprintf("[%d]", rbt.Key)

	if rbt.Color {
		s += "*"

	} else {
		s += "O"
	}
	return s

}
