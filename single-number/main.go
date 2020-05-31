package main

import "fmt"

// 使用map记录出现的次数
func singleNumber(nums []int) int {
	var cnt map[int]int
	cnt = make(map[int]int, len(nums))
	for _, n := range nums {
		if cnt[n] != 0 {
			cnt[n]++
		} else {
			cnt[n] = 1
		}
	}
	for k, v := range cnt {
		if v == 1 {
			return k
		}
	}
	return 0
}

// 使用map构建“只出现一次”的集合
func singleNumber2(nums []int) int {
	var single map[int]int
	single = make(map[int]int, len(nums))
	for _, n := range nums {
		if single[n] != 0 {
			delete(single, n)
		} else {
			single[n] = 1
		}
	}
	for k := range single {
		return k
	}
	return 0
}

// 使用异或
func singleNumber3(nums []int) int {
	var res int
	for _, n := range nums {
		res = res ^ n
	}
	return res
}

func main() {
	fmt.Println(singleNumber3([]int{2, 2, 1}))
}
