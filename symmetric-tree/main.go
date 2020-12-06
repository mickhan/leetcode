package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func serialized(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return serialized(root.Left) + "," + serialized(root.Right) + "," + strconv.Itoa(root.Val)
}

func symmetric(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	symmetric(root.Left)
	symmetric(root.Right)
	root.Left, root.Right = root.Right, root.Left
	return root
}

func isSymmetric(root *TreeNode) bool {
	// 序列化原始二叉树
	tree1 := serialized(root)
	// 镜像翻转
	r := symmetric(root)
	// 序列化翻转后的二叉树
	tree2 := serialized(r)
	return tree1 == tree2
}

func recursion2(l *TreeNode, r *TreeNode) bool {
	if l == r && l == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	return l.Val == r.Val && recursion2(l.Left, r.Right) && recursion2(l.Right, r.Left)
}

func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return recursion2(root.Left, root.Right)
}

func main() {
	nA1 := TreeNode{Val: 1}
	nA2 := TreeNode{Val: 2}
	nA3 := TreeNode{Val: 2}
	nA4 := TreeNode{Val: 3}
	nA5 := TreeNode{Val: 4}
	nA6 := TreeNode{Val: 4}
	nA7 := TreeNode{Val: 3}
	nA1.Left = &nA2
	nA1.Right = &nA3
	nA2.Left = &nA4
	nA2.Right = &nA5
	nA3.Left = &nA6
	nA3.Right = &nA7
	fmt.Println(isSymmetric2(&nA1))
}
