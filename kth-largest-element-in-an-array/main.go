package main

import (
	"fmt"
	"sort"
)

// 先排序，再找
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	return nums[len(nums)-k]
}

// 快速选择
// 基于快速排序

// 交换，使左边小于中间值，右边大于中间值。返回中间值的下标。
func partition(nums []int, l int, r int) int {
	// 选nums[r]作为中间值
	// 前后指针法
	i := l - 1
	for j := l; j < r; j++ {
		if nums[j] <= nums[r] {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	nums[i+1], nums[r] = nums[r], nums[i+1]
	return i + 1
}

// 从l到r，使得l至idx的数小于等于nums[idx]，idx至r的数大于等于nums[idx]
func quickSelect(nums []int, l int, r int, idx int) int {
	q := partition(nums, l, r)
	if q == idx {
		return nums[q]
	} else if q < idx {
		return quickSelect(nums, q+1, r, idx)
	} else {
		return quickSelect(nums, l, q-1, idx)
	}
}

func findKthLargest2(nums []int, k int) int {
	x := quickSelect(nums, 0, len(nums)-1, len(nums)-k)
	return x
}

func main() {
	fmt.Println(findKthLargest2([]int{5, 3, 2, 1, 6, 4}, 2))
	// fmt.Println(findKthLargest2([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	// fmt.Println(findKthLargest2([]int{-1, -1}, 2))
}
