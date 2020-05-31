package main

import "fmt"

func productExceptSelf(nums []int) []int {
	res := make([]int, len(nums))
	// 定义L R数组
	l := make([]int, len(nums))
	r := make([]int, len(nums))
	// 计算i左侧所有元素的乘积
	l[0] = 1 // 左侧没有元素，赋1
	for i := 1; i < len(nums); i++ {
		l[i] = l[i-1] * nums[i-1]
	}
	// 计算i右侧所有元素的乘积
	r[len(nums)-1] = 1 // 右侧没有元素，赋1
	for i := len(nums) - 2; i >= 0; i-- {
		r[i] = r[i+1] * nums[i+1]
	}
	// i位置的乘积是L[i]*R[i]
	for i := 0; i < len(nums); i++ {
		res[i] = l[i] * r[i]
	}
	return res
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
