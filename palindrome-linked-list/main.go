package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// 拷贝链表
func copyList(head *ListNode) *ListNode {
	prehead := ListNode{}
	c := &prehead
	for head != nil {
		node := ListNode{Val: head.Val}
		c.Next = &node
		head, c = head.Next, c.Next
	}
	return prehead.Next
}

// 反转链表
func reverse(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	last := reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return last
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	a := copyList(head)
	b := reverse(head)
	// 判断正反链表是否相等，即回文
	for a != nil {
		if a.Val != b.Val {
			return false
		}
		a, b = a.Next, b.Next
	}
	return true
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func main() {
	l1 := ListNode{Val: 1}
	l2 := ListNode{Val: 1}
	l3 := ListNode{Val: 2}
	l4 := ListNode{Val: 1}
	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	fmt.Println(isPalindrome(&l1))
}
