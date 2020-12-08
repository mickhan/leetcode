package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// root是否符合左<根<右
// 并且左子树都小于根 右子树都大于根
func recursion(root *TreeNode, upper int, lower int) bool {
	if root == nil {
		return true
	}
	if root.Val <= lower || root.Val >= upper {
		return false
	}
	return recursion(root.Left, root.Val, lower) && recursion(root.Right, upper, root.Val)
}

func isValidBST(root *TreeNode) bool {
	return recursion(root, math.MaxInt64, math.MinInt64)
}

func main() {
	n1 := TreeNode{Val: 2}
	n2 := TreeNode{Val: 1}
	n3 := TreeNode{Val: 3}
	n1.Left = &n2
	n1.Right = &n3
	fmt.Println(isValidBST(&n1))

	nA1 := TreeNode{Val: 5}
	nA2 := TreeNode{Val: 4}
	nA3 := TreeNode{Val: 6}
	nA4 := TreeNode{Val: 3}
	nA5 := TreeNode{Val: 7}
	nA1.Left = &nA2
	nA1.Right = &nA3
	nA3.Left = &nA4
	nA3.Right = &nA5
	fmt.Println(isValidBST(&nA1))
}
