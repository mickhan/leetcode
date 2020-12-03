package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// 同时递归两个节点
func recursion(left *Node, right *Node) {
	if left == nil || right == nil {
		return
	}
	left.Next = right
	// 同节点
	recursion(left.Left, left.Right)
	recursion(right.Left, right.Right)
	// 跨节点
	recursion(left.Right, right.Left)
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	recursion(root.Left, root.Right)
	return root
}

func main() {
	n1 := Node{Val: 1}
	n2 := Node{Val: 2}
	n3 := Node{Val: 3}
	n4 := Node{Val: 4}
	n5 := Node{Val: 5}
	n6 := Node{Val: 6}
	n7 := Node{Val: 7}
	n1.Left = &n2
	n1.Right = &n3
	n2.Left = &n4
	n2.Right = &n5
	n3.Left = &n6
	n3.Right = &n7
	connect(&n1)
}
