package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 使用双指针
	a, b := head, head
	// b先移动n+1次
	for n = n + 1; n > 0; n-- {
		// b已经移动到链表末尾，说明n为链表长度，即删除链表第一个节点
		if b == nil {
			return head.Next
		}
		b = b.Next
	}
	// 再同时移动，当b到达结尾的时候，a就指向倒数第n+1个节点
	for b != nil {
		a, b = a.Next, b.Next
	}
	// 删除a.Next，即倒数第n个节点
	a.Next = a.Next.Next
	return head
}

func main() {
	a := ListNode{Val: 1}
	a1 := ListNode{Val: 2}
	a2 := ListNode{Val: 3}
	a3 := ListNode{Val: 4}
	a4 := ListNode{Val: 5}
	a.Next = &a1
	a1.Next = &a2
	a2.Next = &a3
	a3.Next = &a4
	fmt.Println(removeNthFromEnd(&a, 2))
}
