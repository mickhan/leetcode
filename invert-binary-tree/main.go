package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	invertTree(root.Left)
	invertTree(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
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
	fmt.Println(invertTree(&nA1))
}
