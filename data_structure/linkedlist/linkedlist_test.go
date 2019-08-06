package linkedlist

import (
	"testing"
)

func TestHasCycle(t *testing.T) {
	listOfNode := [9]*ListNode{}
	head := &ListNode{0, nil}
	listOfNode[0] = head
	p := head
	for i := 1; i < 9; i++ {
		p = head.Add(i)
		listOfNode[i] = p
	}
	t.Log(head)
	if head.HasCycle() {
		t.Log("Error it should have cycle .")
	}
	p.Next.Next = listOfNode[3]

	if head.HasCycle() {
		t.Log("cycle here", head.CyclePos().Val)
		t.Log(head)
	} else {

		t.Log("Error it should have cycle .")
	}

}

func TestAddNew(t *testing.T) {

	testcases := []struct {
		l1     []int
		l2     []int
		equals bool
	}{
		{
			l1:     []int{},
			l2:     []int{1},
			equals: false,
		},
		{
			l1:     []int{},
			l2:     []int{},
			equals: true,
		},
		{
			l1:     []int{1},
			l2:     []int{1, 1},
			equals: true,
		},
		{
			l1:     []int{1, 2, 3},
			l2:     []int{1, 2, 3, 1},
			equals: true,
		},
		{
			l1:     []int{1, 2},
			l2:     []int{},
			equals: false,
		},
	}
	for i, tcase := range testcases {
		l1 := MakeLinkedList(tcase.l1)
		l2 := MakeLinkedList(tcase.l2)
		l1.Add(1)
		equals := ListNodeEquals(l1, l2)
		if equals == tcase.equals {
			t.Logf("%d/%d PASSED %v ", i+1, len(testcases), l1)
		} else {
			t.Errorf("%d/%d FAILED %v,%v ", i+1, len(testcases), l1, l2)
		}

	}

}

func TestMakeLinkList(t *testing.T) {
	testcases := [][]int{
		{},
		{1},
		{1, 2, 3},
	}
	for _, sl := range testcases {
		l := MakeLinkedList(sl)
		t.Logf("l = [%v]", l)
	}

}
func TestListNodeEquals(t *testing.T) {
	testcases := []struct {
		l1     []int
		l2     []int
		equals bool
	}{
		{
			// 1
			l1:     []int{},
			l2:     []int{},
			equals: true,
		},
		{
			// 2
			l1:     []int{},
			l2:     []int{1},
			equals: false,
		},
		{
			// 3
			l1:     []int{1},
			l2:     []int{1},
			equals: true,
		},
		{
			// 4
			l1:     []int{1},
			l2:     []int{1, 2, 3},
			equals: false,
		},
		{
			// 5
			l1:     []int{1, 2, 3},
			l2:     []int{1, 2, 3},
			equals: true,
		},
		{
			// 6
			l1:     []int{1, 2},
			l2:     []int{},
			equals: false,
		},
	}
	for i, tcase := range testcases {
		l1 := MakeLinkedList(tcase.l1)
		l2 := MakeLinkedList(tcase.l2)
		equals := ListNodeEquals(l1, l2)
		if equals == tcase.equals {
			t.Logf("%d/%d PASSED ", i+1, len(testcases))
		} else {
			t.Errorf("%d/%d FAILED %v,%v ", i+1, len(testcases), l1, l2)
		}

	}

}
