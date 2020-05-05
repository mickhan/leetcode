package main

import "fmt"

var sets [][]int
var gnums []int // 全集
var cur []int   // 当前尝试

func makeOptions(i int, options []int) []int {
	// 剔除第i个，形成新的可选项
	ret := append(options[:i], options[i+1:]...)
	return ret
}

func backtrace(pos int, options []int) {
	// 当前尝试的长度等于全集，说明排列完成，记录到结果中
	if len(cur) == len(gnums) {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		sets = append(sets, tmp)
	}
	// 循环尝试可选项中的所有可能性
	for i := 0; i < len(options); i++ {
		cur = append(cur, options[i])
		opt := make([]int, len(options))
		copy(opt, options)
		backtrace(pos+1, makeOptions(i, opt))
		cur = cur[:len(cur)-1]
	}
}

func permute(nums []int) [][]int {
	sets = make([][]int, 0)
	gnums = nums
	backtrace(0, nums)
	return sets
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
