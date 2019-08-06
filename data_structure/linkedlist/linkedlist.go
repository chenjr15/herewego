package linkedlist

import "fmt"

// ListNode linked list
type ListNode struct {
	Val  int
	Next *ListNode
}

// Add(int) Generate a new element then attach to the tail of list, list may not be nil
func (l *ListNode) Add(x int) *ListNode {
	if l == nil {
		return nil
	}
	for ; l.Next != nil; l = l.Next {
	}
	l.Next = &ListNode{x, nil}
	return l

}

// CyclePos() return the cross position of cycle
func (l *ListNode) CyclePos() (crossAt *ListNode) {
	if l == nil || l.Next == nil {
		return nil
	}
	forward := true
	slow := l
	var p *ListNode
	for fast := l.Next; fast != nil; fast = fast.Next {
		if fast.Next == slow.Next {
			p = slow.Next
			break
		}
		if forward {
			slow = slow.Next
		}
		forward = !forward

	}
	for slow = l; p != nil; {
		if slow == p {
			return slow
		}
		p = p.Next
		slow = slow.Next

	}
	return nil
}

// HasCycle()  Determine if it has a cycle in it
func (l *ListNode) HasCycle() bool {

	return l.CyclePos() != nil
}

// MakeLinkList make a linked list form int slice
func MakeLinkedList(data []int) (head *ListNode) {
	head = new(ListNode)
	p := head
	for _, item := range data {
		p.Next = new(ListNode)
		p = p.Next
		p.Val = item
	}
	p.Next = nil
	head = head.Next
	return
}

// ListNodeEquals judge whether two linked list equals
func ListNodeEquals(l1, l2 *ListNode) bool {
	if l1 == nil || l2 == nil {
		return l1 == l2
	}
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			break
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == l2
}

func (l *ListNode) String() (result string) {
	cycle := l.CyclePos()

	met := false
	for l != nil {
		if l == cycle {
			if met {
				result += "<"
				return
			}
			result += "^"
			met = true

		}
		result += fmt.Sprintf("%v-> ", l.Val)
		l = l.Next
	}
	result += "&"
	return
}
