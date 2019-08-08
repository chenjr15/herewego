package tree

import (
	"container/list"
	"fmt"
)

// Node struct of tree
type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

// New Return a new Node with val
func New(val int) *Node {

	return &Node{
		Val:   val,
		Left:  nil,
		Right: nil,
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
			visit(t, level)
			if t.Left != nil {
				queue.PushBack(t.Left)
				lastLevel++
			}
			if t.Right != nil {
				queue.PushBack(t.Right)
				lastLevel++
			}
		}
		level++

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

// AddLeft go to the end of left and add a new node, then return new node
func (t *Node) AddLeft(val int) *Node {
	if t == nil {
		return nil
	}
	for t.Left != nil {
		t = t.Left
	}
	t.Left = &Node{
		val,
		nil,
		nil,
	}
	return t.Left

}

// AddRight go to the end of right and add a new node, then return new node
func (t *Node) AddRight(val int) *Node {
	if t == nil {
		return nil
	}
	for t.Right != nil {
		t = t.Right
	}
	t.Right = &Node{
		val,
		nil,
		nil,
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
func (t *Node) String() string {
	if t == nil {
		return "[-]"
	}
	return fmt.Sprintf("[%d]", t.Val)

}
