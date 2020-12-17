package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry, v1, v2 int
	prehead := ListNode{}
	idx := &prehead
	// 一直加到两个链表都到末尾
	for l1 != nil || l2 != nil {
		if l1 == nil {
			v1 = 0
		} else {
			v1 = l1.Val
		}
		if l2 == nil {
			v2 = 0
		} else {
			v2 = l2.Val
		}
		node := ListNode{}
		node.Val = (v1 + v2 + carry) % 10 // 本节点储存0-9
		carry = (v1 + v2 + carry) / 10    // 进位
		idx.Next = &node
		idx = idx.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	// 如果仍有进位，添加一个节点
	if carry != 0 {
		node := ListNode{Val: carry}
		idx.Next = &node
	}
	return prehead.Next
}

func main() {
	a := ListNode{Val: 2}
	a1 := ListNode{Val: 4}
	a2 := ListNode{Val: 3}
	a.Next = &a1
	a1.Next = &a2
	b := ListNode{Val: 5}
	b1 := ListNode{Val: 6}
	b2 := ListNode{Val: 4}
	b.Next = &b1
	b1.Next = &b2
	h := addTwoNumbers(&a, &b)
	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}
