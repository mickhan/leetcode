package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return recursion(nums)
}

func recursion(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	// 找到最大值
	maxNums := nums[0]
	var idx int
	for i := 0; i < len(nums); i++ {
		if nums[i] >= maxNums {
			maxNums = nums[i]
			idx = i
		}
	}
	// 创建当前节点
	node := TreeNode{Val: maxNums}
	// fmt.Println(nums, idx, node.Val)
	// 数组拆分成两半
	// 左
	if idx > 0 {
		l := make([]int, idx)
		copy(l, nums[:idx])
		node.Left = recursion(l)
	}
	// 右
	if len(nums)-idx > 0 {
		r := make([]int, len(nums)-idx-1)
		copy(r, nums[idx+1:])
		node.Right = recursion(r)
	}
	return &node
}

func main() {
	recursion([]int{3, 2, 1, 6, 0, 5})
}
