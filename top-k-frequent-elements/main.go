package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	// 统计出现次数
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	// 对统计结果排序
	type kv struct {
		num  int
		freq int
	}
	var arr []kv
	for k, v := range freq {
		arr = append(arr, kv{k, v})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].freq > arr[j].freq
	})
	var res []int
	for i := 0; i < k; i++ {
		res = append(res, arr[i].num)
	}
	return res
}

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2))
}
