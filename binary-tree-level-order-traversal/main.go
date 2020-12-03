package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	var q []*TreeNode
	q = append(q, root)
	for len(q) != 0 {
		var levelRes []int
		levelLen := len(q)
		for i := 0; i < levelLen; i++ {
			n := q[0]
			levelRes = append(levelRes, n.Val)
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}
			q = q[1:]
		}
		tmp := make([]int, levelLen)
		copy(tmp, levelRes)
		res = append(res, tmp)
	}
	return res
}

func main() {
	n1 := TreeNode{Val: 3}
	n2 := TreeNode{Val: 9}
	n3 := TreeNode{Val: 20}
	n4 := TreeNode{Val: 15}
	n5 := TreeNode{Val: 7}
	n1.Left = &n2
	n1.Right = &n3
	n3.Left = &n4
	n3.Right = &n5

	fmt.Println(levelOrder(&n1))
}
