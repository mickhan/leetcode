package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	fmt.Println(inorder, postorder)
	var root TreeNode
	if len(postorder) == 0 {
		return nil
	}
	// 根是postorder最后一个
	root.Val = postorder[len(postorder)-1]
	if len(postorder) == 1 {
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
	// 根据根的位置拆分成左右两部分，仍然保持中序/后序
	iLeft := inorder[:pos]
	iRight := inorder[pos+1:]
	pLeft := postorder[:len(iLeft)]
	pRight := postorder[len(iLeft) : len(postorder)-1]
	root.Left = buildTree(iLeft, pLeft)
	root.Right = buildTree(iRight, pRight)
	return &root
}

func main() {
	fmt.Println(buildTree([]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}))
	// fmt.Println(buildTree([]int{}, []int{}))
}
