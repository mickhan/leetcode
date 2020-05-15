package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	inorder(root.Left, res)
	*res = append(*res, root.Val)
	inorder(root.Right, res)
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	res = make([]int, 0)
	inorder(root, &res)
	return res
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
