package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recordParents(root *TreeNode, parents map[*TreeNode]*TreeNode) {
	if root.Left != nil {
		parents[root.Left] = root
		recordParents(root.Left, parents)
	}
	if root.Right != nil {
		parents[root.Right] = root
		recordParents(root.Right, parents)
	}
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// 构建一个node与其父节点映射关系
	parents := make(map[*TreeNode]*TreeNode, 0)
	recordParents(root, parents)
	// 依次在其中找到p、q的祖先
	var pA, qA []*TreeNode
	var tmp *TreeNode
	tmp = p
	for tmp != nil {
		pA = append(pA, tmp)
		tmp = parents[tmp]
	}
	tmp = q
	for tmp != nil {
		qA = append(qA, tmp)
		tmp = parents[tmp]
	}
	// 判断两组祖先的交集，得到最小公共祖先
	for _, i := range pA {
		for _, j := range qA {
			if i == j {
				return i
			}
		}
	}
	return nil
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
	fmt.Println(lowestCommonAncestor(&nA1, &nA3, &nA4))
}
