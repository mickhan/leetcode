package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var sum int

// 反向中序遍历BST，使得遍历结果从大到小排列
func recursion(root *TreeNode) {
	if root == nil {
		return
	}
	recursion(root.Right)
	sum = sum + root.Val
	root.Val = sum
	recursion(root.Left)
}

func convertBST(root *TreeNode) *TreeNode {
	sum = 0
	recursion(root)
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	n1 := TreeNode{Val: 3}
	n2 := TreeNode{Val: 2}
	n3 := TreeNode{Val: 4}
	n4 := TreeNode{Val: 1}
	n1.Left = &n2
	n1.Right = &n3
	n2.Left = &n4

	printTree(convertBST(&n1))
}
