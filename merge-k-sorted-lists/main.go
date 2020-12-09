package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 合并两个链表
func mergeList(a, b *ListNode) *ListNode {
	prehead := ListNode{}
	t := &prehead
	for a != nil && b != nil {
		if a.Val <= b.Val {
			t.Next = a
			t, a = t.Next, a.Next
		} else {
			t.Next = b
			t, b = t.Next, b.Next
		}
	}
	if a == nil {
		t.Next = b
	} else {
		t.Next = a
	}
	return prehead.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	var h *ListNode
	// 每次合并两个链表
	for _, l := range lists {
		h = mergeList(h, l)
	}
	return h
}

func main() {
	a := ListNode{Val: 1}
	a1 := ListNode{Val: 4}
	a2 := ListNode{Val: 5}
	a.Next = &a1
	a1.Next = &a2
	b := ListNode{Val: 1}
	b1 := ListNode{Val: 3}
	b2 := ListNode{Val: 4}
	b.Next = &b1
	b1.Next = &b2
	c := ListNode{Val: 2}
	c1 := ListNode{Val: 6}
	c.Next = &c1

	h := mergeKLists([]*ListNode{&a, &b, &c})
	for h != nil {
		fmt.Printf("%d ", h.Val)
		h = h.Next
	}
}
