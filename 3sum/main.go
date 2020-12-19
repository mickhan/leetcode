package main

import (
	"fmt"
	"sort"
	"strings"
)

// 暴力解
// O(n^3)会超时
func threeSum(nums []int) [][]int {
	mem := make(map[string]bool)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					tmp := []int{nums[i], nums[j], nums[k]}
					sort.Ints(tmp)
					key := strings.Trim(strings.Replace(fmt.Sprint(tmp), " ", ",", -1), "[]")
					if _, ok := mem[key]; !ok {
						// 不存在，说明没有重复
						res = append(res, []int{nums[i], nums[j], nums[k]})
						mem[key] = true
					}
				}
			}
		}
	}
	return res
}

// 排序+双指针
func threeSum2(nums []int) [][]int {
	sort.Ints(nums)
	fmt.Println(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for {
			for j > i+1 && j < len(nums) && nums[j] == nums[j-1] {
				j++
			}
			for k > j && k < len(nums)-1 && nums[k] == nums[k+1] {
				k--
			}
			if j >= k {
				break
			}
			fmt.Println(nums[i], nums[j], nums[k])
			if nums[i]+nums[j]+nums[k] == 0 {
				res = append(res, []int{nums[i], nums[j], nums[k]})
				j++ // 要求结果不重复的，只变一个的话不可能和为0，需要两个都变
				k--
			} else if nums[i]+nums[j]+nums[k] > 0 { // 总体大于0，需要缩小。因为是递增的，减小k。
				k--
			} else { // 总体小于0，需要增大。
				j++
			}
		}
	}
	return res
}

func main() {
	fmt.Println(threeSum2([]int{-1, 0, 1, 2, -1, -4}))
}
