package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var nth int
var kth int

// 中序遍历BST，找到第一个
func recursion(root *TreeNode, k int) {
	if root == nil {
		return
	}
	recursion(root.Left, k)
	nth++
	if nth == k {
		kth = root.Val
	}
	recursion(root.Right, k)
}

func kthSmallest(root *TreeNode, k int) int {
	// 重置kth, nth
	kth = 0
	nth = 0

	recursion(root, k)
	return kth
}

func main() {
	n1 := TreeNode{Val: 3}
	n2 := TreeNode{Val: 1}
	n3 := TreeNode{Val: 4}
	n4 := TreeNode{Val: 2}
	n1.Left = &n2
	n1.Right = &n3
	n2.Right = &n4

	fmt.Println(kthSmallest(&n1, 1))
}
