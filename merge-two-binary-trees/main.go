package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	node := TreeNode{}
	if t1 == nil && t2 == nil {
		return nil
	} else if t1 != nil && t2 == nil {
		node.Val = t1.Val
		node.Left = mergeTrees(t1.Left, nil)
		node.Right = mergeTrees(t1.Right, nil)
	} else if t1 == nil && t2 != nil {
		node.Val = t2.Val
		node.Left = mergeTrees(nil, t2.Left)
		node.Right = mergeTrees(nil, t2.Right)
	} else {
		node.Val = t1.Val + t2.Val
		node.Left = mergeTrees(t1.Left, t2.Left)
		node.Right = mergeTrees(t1.Right, t2.Right)
	}
	return &node
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

	nB1 := TreeNode{
		Val: 2,
	}
	nB2 := TreeNode{
		Val: 1,
	}
	nB3 := TreeNode{
		Val: 3,
	}
	nB4 := TreeNode{
		Val: 4,
	}
	nB5 := TreeNode{
		Val: 7,
	}
	nB1.Left = &nB2
	nB1.Right = &nB3
	nB2.Right = &nB4
	nB3.Right = &nB5

	mergeTrees(&nA1, &nB1)
}
