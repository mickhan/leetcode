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

// 递归
func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	} else {
		n := reverseList(head.Next)
		head.Next.Next = head // head.Next 就是n的最后一个元素，所以不需要再遍历一遍找了
		head.Next = nil
		return n
	}
}

// 迭代
func reverseList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead, node *ListNode
	// 每次从旧链表的头部拆下一个节点，插入到新链表的头部，即可反向
	for head != nil {
		node = head
		head = head.Next
		node.Next = newHead
		newHead = node
	}
	return newHead
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
	printList(reverseList3(&n1))
}
