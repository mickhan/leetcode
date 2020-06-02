package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var root TreeNode
	if len(preorder) == 0 {
		return nil
	}
	// 根是preorder[0]
	root.Val = preorder[0]
	if len(preorder) == 1 {
		return &root
	}
	var pos int
	// 找到根在inorder中的位置
	for i := range inorder {
		if root.Val == inorder[i] {
			pos = i
			break
		}
	}
	iLeft := inorder[:pos]
	iRight := inorder[pos+1:]
	pLeft := preorder[1 : len(iLeft)+1]
	pRight := preorder[len(iLeft)+1:]
	root.Left = buildTree(pLeft, iLeft)
	root.Right = buildTree(pRight, iRight)
	return &root
}

func main() {
	// fmt.Println(buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}))
	fmt.Println(buildTree([]int{}, []int{}))
}
