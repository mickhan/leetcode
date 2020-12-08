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

func recursion(root *TreeNode, seq map[*TreeNode]int) int {
	if root == nil {
		seq[root] = math.MinInt32
		return 0
	}
	l := recursion(root.Left, seq)
	r := recursion(root.Right, seq)
	// 经过root节点的路径长度（左右只能选一条，也可以都不选）
	arr := []int{root.Val, l + root.Val, r + root.Val}
	mx := arr[0]
	for _, v := range arr {
		if v >= mx {
			mx = v
		}
	}
	// 记录以root为根节点的路径长度（可以左右都选）
	if l+r+root.Val > mx {
		seq[root] = l + r + root.Val
	} else {
		seq[root] = mx
	}
	return mx
}

func maxPathSum(root *TreeNode) int {
	var seq map[*TreeNode]int
	seq = make(map[*TreeNode]int)
	recursion(root, seq)
	// fmt.Println(seq)
	// for k, v := range seq {
	// 	if k != nil {
	// 		fmt.Println(k.Val, v)
	// 	}
	// }
	mx := math.MinInt32
	for _, s := range seq {
		if s > mx {
			mx = s
		}
	}
	return mx
}

func main() {
	nA1 := TreeNode{Val: -10}
	nA2 := TreeNode{Val: 9}
	nA3 := TreeNode{Val: 20}
	nA4 := TreeNode{Val: 15}
	nA5 := TreeNode{Val: 7}
	nA1.Left = &nA2
	nA1.Right = &nA3
	nA3.Left = &nA4
	nA3.Right = &nA5
	fmt.Println(maxPathSum(&nA1))

	// n1 := TreeNode{Val: 2}
	// n2 := TreeNode{Val: -1}
	// n1.Left = &n2
	// fmt.Println(maxPathSum(&n1))

	n1 := TreeNode{Val: 5}
	n2 := TreeNode{Val: 4}
	n3 := TreeNode{Val: 8}
	n4 := TreeNode{Val: 11}
	n5 := TreeNode{Val: 13}
	n6 := TreeNode{Val: 4}
	n7 := TreeNode{Val: 7}
	n8 := TreeNode{Val: 2}
	n9 := TreeNode{Val: 1}
	n1.Left = &n2
	n1.Right = &n3
	n2.Left = &n4
	n3.Left = &n5
	n3.Right = &n6
	n4.Left = &n7
	n4.Right = &n8
	n6.Right = &n9
	fmt.Println(maxPathSum(&n1))
}
