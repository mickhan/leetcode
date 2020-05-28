package main

import (
	"fmt"
	"sort"
)

func cSum(cur []int) int {
	var s int
	for _, c := range cur {
		s += c
	}
	return s
}

func comb(candidates []int, limit int, target int, cur []int, res *[][]int) {
	// 遍历candidates中的每个数字
	// 如果等于target，则说明是有效排列，加入到res
	// 如果大于target，说明后面的数字都无效，可以返回（candidates需要升序排列）
	// 如果小于target，说明还可以再添加一个数字到排列中
	for idx, c := range candidates[limit:] {
		cur = append(cur, c)
		s := cSum(cur)
		if s == target {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			*res = append(*res, tmp)
		} else if s > target {
			return
		} else {
			comb(candidates, limit+idx, target, cur, res)
		}
		cur = cur[:len(cur)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var cur []int
	sort.Ints(candidates) // 解法要求升序排列，处理一下
	comb(candidates, 0, target, cur, &res)
	return res
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
	fmt.Println(combinationSum([]int{2, 3, 5}, 8))
	fmt.Println(combinationSum([]int{8, 7, 4, 3}, 11))
}
