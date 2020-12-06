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

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// 计算左右子树的层数
	var nL, nR int
	tmp := root
	for tmp != nil {
		tmp = tmp.Left
		nL++
	}
	tmp = root
	for tmp != nil {
		tmp = tmp.Right
		nR++
	}
	// 如果nL等于nR，则是满二叉树，节点数可以计算得出
	if nL == nR {
		return int(math.Pow(2, float64(nL))) - 1
	}
	// 否则是完全二叉树，递归计算
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func main() {
	nA1 := TreeNode{Val: 1}
	nA2 := TreeNode{Val: 2}
	nA3 := TreeNode{Val: 3}
	nA4 := TreeNode{Val: 4}
	nA5 := TreeNode{Val: 5}
	nA6 := TreeNode{Val: 6}
	nA1.Left = &nA2
	nA1.Right = &nA3
	nA2.Left = &nA4
	nA2.Right = &nA5
	nA3.Left = &nA6
	fmt.Println(countNodes(&nA1))
}
