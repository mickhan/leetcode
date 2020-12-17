package main

import "fmt"

// 用map记录哪些数字出现过
func findDisappearedNumbers(nums []int) []int {
	var res []int
	present := make(map[int]bool)
	for _, v := range nums {
		present[v] = true
	}
	// fmt.Println(present)
	for i := 1; i < len(nums)+1; i++ {
		if v := present[i]; !v {
			res = append(res, i)
		}
	}
	return res
}

// 把nums[i]-1位置的数置为相反数（仅一次）。所有出现过的都会变成负数，未出现的仍然为正数。
func findDisappearedNumbers2(nums []int) []int {
	var res []int
	for _, value := range nums {
		if value < 0 {
			value = -value
		}
		if nums[value-1] < 0 {
			continue
		}
		nums[value-1] = -nums[value-1]
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {
	fmt.Println(findDisappearedNumbers2([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
