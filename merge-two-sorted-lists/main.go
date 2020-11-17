package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var tail *ListNode
	prehead := &ListNode{}
	tail = prehead
	for l1 != nil && l2 != nil {
		// 把较小的追加到tail
		if l1.Val <= l2.Val {
			tail.Next = l1
			l1 = l1.Next
		} else {
			tail.Next = l2
			l2 = l2.Next
		}
		// 移动tail
		tail = tail.Next
	}
	if l1 == nil {
		tail.Next = l2
	}
	if l2 == nil {
		tail.Next = l1
	}
	return prehead.Next
}

func main() {

}
