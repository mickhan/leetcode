package main

import (
	"fmt"
	"math"
)

// 设重复的数字为k
// 如果k<mid，那么nums中小于mid的个数，就比mid大
func cnt(nums *[]int, mid float64) bool {
	var c int
	for _, n := range *nums {
		if float64(n) <= mid {
			c++
		}
	}
	return c > int(math.Floor(mid))
}

func findDuplicate(nums []int) int {
	var upper, lower int
	var mid float64
	lower = 1
	upper = len(nums) - 1
	mid = float64(upper) / 2
	for upper != lower {
		// 如果小于mid的更多，说明重复的数low<k<mid
		if cnt(&nums, mid) {
			upper = int(math.Floor(mid))
		} else {
			lower = int(math.Ceil(mid))
		}
		mid = float64(upper-lower)/2 + float64(lower)
	}
	return upper
}

func main() {
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 2}))
	fmt.Println(findDuplicate([]int{3, 1, 3, 4, 2}))
	fmt.Println(findDuplicate([]int{1, 1, 2}))
	fmt.Println(findDuplicate([]int{1, 3, 4, 2, 1}))
}
