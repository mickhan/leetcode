package main

import "fmt"


// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func sortList(head *ListNode) *ListNode {
    var cur, prev, next, tail, min *ListNode
    var n int
    tmp := head
    if head == nil {
        return head
    }
    // 计算链表长度
    for {
        n++
        tail = tmp
        tmp = tmp.Next
        if tmp == nil {
            break 
        }
    }
    if n == 1 {
        return head
    }
    // fmt.Println(n)
    // 遍历链表，将最小的数放到链表末尾
    for i := n; i > 0; i-- {
        cur = head
        min = cur
        var tmpPrev *ListNode
        for j := 0; j < i; j++ {
            // fmt.Printf("i:%d j:%d\n", i, j)
            // 发现一个较小的数，记住它，以及它前后的元素
            if cur.Val < min.Val {
                prev = tmpPrev
                min = cur
                next = cur.Next
            }
            tmpPrev = cur
            cur = cur.Next
        }
        // 把找到的最小数，移动到链表末尾
        // 如果tail就是最小数，无需操作
        if tail == min {
            continue
        }
        // 如果head就是最小数，先记录head，避免链表丢失
        if head == min {
            head = min.Next
        } else {
            prev.Next = next
        }
        tail.Next = min
        min.Next = nil
        tail = min
    }
    return head
}

func sortList2(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    // 使用快慢指针找到中点
    slow, fast := head, head.Next
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    // 归并排序
    right := sortList(slow.Next)
    slow.Next = nil
    left := sortList(head)
    prehead := &ListNode{}
    prehead.Next = head
    // 合并结果
    return mergeList(prehead, left, right)
}

func mergeList(prehead *ListNode, l1 *ListNode, l2 *ListNode) *ListNode {
    tail := prehead
    // 合并直到l1/l2为空
    for l1 != nil && l2 != nil{
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
    // 把不为空的链表接在末尾
    if l1 == nil {
        tail.Next = l2
    }
    if l2 == nil {
        tail.Next = l1
    }
    return prehead.Next
}

func main() {
    n1 := &ListNode{ Val: 1 }
    n2 := &ListNode{ Val: 5 }
    n3 := &ListNode{ Val: 3 }
    n4 := &ListNode{ Val: 4 }
    n5 := &ListNode{ Val: 0 }
    n1.Next = n2
    n2.Next = n3
    n3.Next = n4
    n4.Next = n5

    h := sortList2(n1)

    for {
        if h == nil {
           break 
        }
        fmt.Println(h.Val)
        h = h.Next
    }
}