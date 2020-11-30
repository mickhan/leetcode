package main

import "fmt"

var sets [][]int
var gnums []int // 全集
var cur []int   // 当前尝试

func makeOptions(i int, options []int) []int {
	// 避免每层递归的options被修改，拷贝一份新的出来
	opt := make([]int, len(options))
	copy(opt, options)
	// 剔除第i个，形成新的可选项
	ret := append(opt[:i], opt[i+1:]...)
	return ret
}

func backtrace(options []int) {
	// 当前尝试的长度等于全集，说明排列完成，记录到结果中
	if len(cur) == len(gnums) {
		// 直接append是cur的引用，会被修改，拷贝一份出来
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		sets = append(sets, tmp)
	}
	// 循环尝试可选项中的所有可能性
	for i := 0; i < len(options); i++ {
		cur = append(cur, options[i])
		backtrace(makeOptions(i, options))
		cur = cur[:len(cur)-1]
	}
}

func permute(nums []int) [][]int {
	sets = make([][]int, 0)
	gnums = nums
	backtrace(nums)
	return sets
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
