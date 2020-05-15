package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var gRes []int

func inorder(root *TreeNode) {
	if root == nil {
		return
	}
	inorder(root.Left)
	gRes = append(gRes, root.Val)
	inorder(root.Right)
}

func inorderTraversal(root *TreeNode) []int {
	gRes = make([]int, 0)
	inorder(root)
	return gRes
}

func main() {
	nA1 := TreeNode{
		Val: 1,
	}
	nA2 := TreeNode{
		Val: 3,
	}
	nA3 := TreeNode{
		Val: 2,
	}
	nA4 := TreeNode{
		Val: 5,
	}
	nA1.Left = &nA2
	nA1.Right = &nA3
	nA2.Left = &nA4
	fmt.Println(inorderTraversal(&nA1))
}
