package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	x, y := head, head
	for y != nil {
		if x.Val != y.Val {
			x.Next = y
			x = y
		}
		y = y.Next
	}
	x.Next = nil
	return head
}

func main() {
	a := ListNode{Val: 1}
	a1 := ListNode{Val: 1}
	a2 := ListNode{Val: 2}
	a3 := ListNode{Val: 3}
	a4 := ListNode{Val: 3}
	a.Next = &a1
	a1.Next = &a2
	a2.Next = &a3
	a3.Next = &a4
	h := deleteDuplicates(&a)
	for h != nil {
		fmt.Printf("%d ", h.Val)
		h = h.Next
	}
}
