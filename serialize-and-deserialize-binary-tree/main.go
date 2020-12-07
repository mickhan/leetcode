package main

import (
	"fmt"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + "," + this.serialize(root.Right)
}

func recursion(data []string) (*TreeNode, int) {
	if len(data) == 0 {
		return nil, 0
	}
	t := data[0]
	data = data[1:]
	if t == "#" {
		return nil, 1
	}
	node := TreeNode{}
	node.Val, _ = strconv.Atoi(t)
	var lenL, lenR int
	fmt.Println(data)
	node.Left, lenL = recursion(data)
	data = data[lenL:]
	fmt.Println(data)
	node.Right, lenR = recursion(data)
	fmt.Println(node.Val, lenL, lenR)
	return &node, lenL + lenR + 1
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	r, _ := recursion(strings.Split(data, ","))
	return r
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	nA1 := TreeNode{Val: 1}
	nA2 := TreeNode{Val: 2}
	nA3 := TreeNode{Val: 3}
	nA4 := TreeNode{Val: 4}
	nA5 := TreeNode{Val: 5}
	nA1.Left = &nA2
	nA1.Right = &nA3
	nA3.Left = &nA4
	nA3.Right = &nA5

	ser := Constructor()
	deser := Constructor()
	data := ser.serialize(&nA1)
	// fmt.Println(data)
	// fmt.Println()
	ans := deser.deserialize(data)
	// fmt.Println()
	printTree(ans)
}
