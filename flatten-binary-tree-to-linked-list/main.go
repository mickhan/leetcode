package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getTail(h *TreeNode) *TreeNode {
	for h.Right != nil {
		h = h.Right
	}
	return h
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	tmp := root.Right
	// 把左节点移动到右侧
	root.Right = root.Left
	root.Left = nil
	// 把原来的右节点追加到尾部
	tail := getTail(root)
	tail.Right = tmp
}

func main() {
	nA1 := TreeNode{
		Val: 1,
	}
	nA2 := TreeNode{
		Val: 2,
	}
	nA3 := TreeNode{
		Val: 3,
	}
	nA4 := TreeNode{
		Val: 4,
	}
	nA5 := TreeNode{
		Val: 5,
	}
	nA6 := TreeNode{
		Val: 6,
	}
	nA1.Left = &nA2
	nA1.Right = &nA5
	nA2.Left = &nA3
	nA2.Right = &nA4
	nA5.Right = &nA6
	flatten(&nA1)
}
