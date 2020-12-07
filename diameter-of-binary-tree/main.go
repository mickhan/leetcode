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

// root为根的树的最大深度是多少
func maxDepth(root *TreeNode, depth map[*TreeNode]int) int {
	if root == nil {
		depth[root] = -1
		return -1
	}
	d := int(math.Max(float64(maxDepth(root.Left, depth)), float64(maxDepth(root.Right, depth)))) + 1
	depth[root] = d
	return d
}

// 穿过root的最长路径是多少
func maxRoute(root *TreeNode, depth map[*TreeNode]int, route map[*TreeNode]int) {
	if root == nil {
		route[root] = 0
		return
	}
	route[root] = depth[root.Left] + depth[root.Right] + 2
	maxRoute(root.Left, depth, route)
	maxRoute(root.Right, depth, route)
}

func diameterOfBinaryTree(root *TreeNode) int {
	var depth map[*TreeNode]int
	depth = make(map[*TreeNode]int)
	var route map[*TreeNode]int
	route = make(map[*TreeNode]int)
	maxDepth(root, depth)
	fmt.Println(depth)
	maxRoute(root, depth, route)
	fmt.Println(route)
	mr := 0
	for _, r := range route {
		if r > mr {
			mr = r
		}
	}
	return mr
}

func main() {
	nA1 := TreeNode{Val: 1}
	nA2 := TreeNode{Val: 2}
	nA3 := TreeNode{Val: 3}
	nA4 := TreeNode{Val: 4}
	nA5 := TreeNode{Val: 5}
	nA1.Left = &nA2
	nA1.Right = &nA3
	nA2.Left = &nA4
	nA2.Right = &nA5
	fmt.Println(diameterOfBinaryTree(&nA1))
}
