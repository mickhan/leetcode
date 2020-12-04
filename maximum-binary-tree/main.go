package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return recursion(nums, 0, len(nums))
}

func recursion(nums []int, start int, end int) *TreeNode {
	if start >= end {
		return nil
	}
	// 找到最大值
	maxNums := nums[start]
	var idx int
	for i := start; i < end; i++ {
		if nums[i] >= maxNums {
			maxNums = nums[i]
			idx = i
		}
	}
	// 创建当前节点
	node := TreeNode{Val: maxNums}
	// fmt.Println(nums, start, end, idx, node.Val)
	// 数组拆分成两半
	node.Left = recursion(nums, start, idx)
	node.Right = recursion(nums, idx+1, end)
	return &node
}

func main() {
	constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
}
