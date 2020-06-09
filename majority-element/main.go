package main

import (
	"fmt"
	"sort"
)

func majorityElement(nums []int) int {
	var major, max int
	cnt := make(map[int]int, 0)
	for _, i := range nums {
		// 记录每个数出现次数
		cnt[i]++
		// 维护最大值
		if cnt[i] >= max {
			max = cnt[i]
			major = i
		}
	}
	return major
}

func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func main() {
	fmt.Println(majorityElement2([]int{3, 2, 3}))
}
