package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	// 处理为空的情况
	if l1 == nil {
		head = l2
		return head
	}
	if l2 == nil {
		head = l1
		return head
	}
	// 决定head指向哪
	if l1.Val <= l2.Val {
		head = l1
		l1 = l1.Next
	} else {
		head = l2
		l2 = l2.Next
	}
	tail = head
	for {
		// 遍历完了l1/l2，把剩余的追加到tail，退出
		if l1 == nil {
			tail.Next = l2
			break
		}
		if l2 == nil {
			tail.Next = l1
			break
		}
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
	return head
}

func main() {

}
