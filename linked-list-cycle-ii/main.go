package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head.Next
	for slow != fast {
		if fast == nil || fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	// slow是快慢指针相遇点
	// head到入环点，相遇点到入环点，两者有数学关系
	slow = slow.Next
	h := head
	for h != slow {
		h = h.Next
		slow = slow.Next
	}
	return h
}

func main() {
	a := ListNode{Val: 1}
	a1 := ListNode{Val: 2}
	a.Next = &a1
	a1.Next = &a
	fmt.Println(detectCycle(&a))
}
