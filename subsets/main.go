package main

import "fmt"

var sets [][]int
var sub []int
var gnums []int

func backtrace(pos int) {
	var n int
	c := make([]int, len(sub))
	copy(c, sub)
	sets = append(sets, c)
	for n = pos; n < len(gnums); n++ {
		sub = append(sub, gnums[n])
		backtrace(n + 1)
		sub = sub[:len(sub)-1]
	}
}

func subsets(nums []int) [][]int {
	sets = make([][]int, 0)
	gnums = nums
	backtrace(0)
	return sets
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}
