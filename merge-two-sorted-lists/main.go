package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	if l1 == nil {
		head = l2
		return head
	}
	if l2 == nil {
		head = l1
		return head
	}
	if l1.Val <= l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	tail = head
	for {
		if l1 == nil {
			tail.Next = l2
			break
		}
		if l2 == nil {
			tail.Next = l1
			break
		}
		if l1.Val <= l2.Val {
			tail.Next = l1
			tail = tail.Next
			l1 = l1.Next
		} else {
			tail.Next = l2
			tail = tail.Next
			l2 = l2.Next
		}
	}
	return head
}

func main() {

}
