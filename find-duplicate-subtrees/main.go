package main

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursion(root *TreeNode, subtreeRoot map[string]*TreeNode, subtreeNum map[string]int) string {
	if root == nil {
		return ""
	}
	l := recursion(root.Left, subtreeRoot, subtreeNum)
	r := recursion(root.Right, subtreeRoot, subtreeNum)
	st := strconv.Itoa(root.Val)
	if l != "" {
		st = l + "-" + st
	}
	if r != "" {
		st = st + "+" + r
	}
	subtreeRoot[st] = root
	subtreeNum[st]++
	return st
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	var res []*TreeNode
	var subtreeRoot map[string]*TreeNode
	var subtreeNum map[string]int
	subtreeRoot = make(map[string]*TreeNode)
	subtreeNum = make(map[string]int)
	recursion(root, subtreeRoot, subtreeNum)
	for k, v := range subtreeNum {
		if v > 1 {
			res = append(res, subtreeRoot[k])
		}
	}
	// fmt.Println(subtreeRoot)
	// fmt.Println(subtreeNum)
	return res
}

func main() {
	nA1 := TreeNode{Val: 1}
	nA2 := TreeNode{Val: 2}
	nA3 := TreeNode{Val: 3}
	nA4 := TreeNode{Val: 4}
	nA5 := TreeNode{Val: 2}
	nA6 := TreeNode{Val: 4}
	nA7 := TreeNode{Val: 4}
	nA1.Left = &nA2
	nA1.Right = &nA3
	nA2.Left = &nA4
	nA3.Left = &nA5
	nA3.Right = &nA6
	nA5.Left = &nA7
	// findDuplicateSubtrees(&nA1)

	n1 := TreeNode{Val: 0}
	n2 := TreeNode{Val: 0}
	n3 := TreeNode{Val: 0}
	n4 := TreeNode{Val: 0}
	n5 := TreeNode{Val: 0}
	n6 := TreeNode{Val: 0}
	n1.Left = &n2
	n1.Right = &n3
	n2.Left = &n4
	n3.Right = &n5
	n5.Right = &n6
	findDuplicateSubtrees(&n1)
}
