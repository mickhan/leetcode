package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func appendList(head *ListNode, node *ListNode) *ListNode {
	var tail *ListNode
	tail = head
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = node
	return head
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else {
		n := reverseList(head.Next)
		head.Next = nil
		return appendList(n, head)
	}
}

func printList(head *ListNode) {
	if head == nil {
		return
	}
	fmt.Println(head.Val)
	printList(head.Next)
}

func main() {
	n1 := ListNode{
		Val: 1,
	}
	n2 := ListNode{
		Val: 2,
	}
	n3 := ListNode{
		Val: 3,
	}
	n1.Next = &n2
	n2.Next = &n3
	printList(reverseList(&n1))
}
