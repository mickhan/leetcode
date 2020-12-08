package main

import (
	"fmt"
	"math"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 直接对比两个链表，O(m*n)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ha := headA
	hb := headB
	for ha != nil {
		for hb != nil {
			fmt.Println(ha.Val, hb.Val)
			if ha == hb {
				return ha
			}
			hb = hb.Next
		}
		ha = ha.Next
		hb = headB
	}
	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	ha := headA
	hb := headB
	var c, lenA, lenB int
	var fa, fb bool
	lenA, lenB = math.MaxInt32, math.MaxInt32
	for c <= lenA+lenB {
		fmt.Println(ha, hb)
		if ha == hb {
			return ha
		}
		ha = ha.Next
		hb = hb.Next
		c++
		if ha == nil && !fa {
			lenA = c
			ha = headB
			fa = true
		}
		if hb == nil && !fb {
			lenB = c
			hb = headA
			fb = true
		}
	}
	return nil
}

func getIntersectionNode3(headA, headB *ListNode) *ListNode {
	ha, hb := headA, headB
	// 如果相交，比如 A[1, 2, 3]，B[2, 3]
	// 接上之后变成
	// [1, 2, 3][2, 3]
	// [2, 3][1, 2, 3]
	// 尾部对齐了

	// 如果没有相交，比如 A[1, 2, 3]，B[4, 5]
	// 接上之后变成
	// [1, 2, 3][4, 5]->nil
	// [4, 5][1, 2, 3]->nil
	// 最终nil仍然会对齐
	for ha != hb {
		if ha == nil {
			ha = headB
		} else {
			ha = ha.Next
		}
		if hb == nil {
			hb = headA
		} else {
			hb = hb.Next
		}
	}
	return ha
}

func main() {
	a := ListNode{Val: 4}
	a1 := ListNode{Val: 1}
	a.Next = &a1
	b := ListNode{Val: 5}
	b1 := ListNode{Val: 0}
	b2 := ListNode{Val: 1}
	b.Next = &b1
	b1.Next = &b2
	c := ListNode{Val: 8}
	c1 := ListNode{Val: 4}
	c2 := ListNode{Val: 5}
	c.Next = &c1
	c1.Next = &c2
	a1.Next = &c
	b2.Next = &c
	fmt.Println(getIntersectionNode3(&a, &b) == &c)

	x := ListNode{Val: 1}
	y := ListNode{Val: 3}
	x.Next = &y
	fmt.Println(getIntersectionNode3(&x, &y) == &y)
}
