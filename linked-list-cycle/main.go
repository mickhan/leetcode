package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	var visit map[*ListNode]bool
	visit = make(map[*ListNode]bool)
	for head == nil {
		if _, ok := visit[head]; ok {
			return true
		}
		visit[head] = true
		head = head.Next
	}
	return false
}

// 快慢指针
func hasCycle2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow, fast := head, head.Next
	for fast != slow {
		if fast == nil || fast.Next == nil {
			return false
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return true
}

func main() {
	a := ListNode{Val: 1}
	a1 := ListNode{Val: 2}
	a.Next = &a1
	a1.Next = &a
	fmt.Println(hasCycle2(&a))
}
