package main

import (
	"fmt"
)

func countBits(num int) []int {
	var res []int
	res = make([]int, num+1)
	res[0] = 0
	for i := 1; i <= num; i++ {
		if i%2 == 1 {
			res[i] = res[i-1] + 1
		} else {
			res[i] = res[i/2]
		}
	}
	return res
}

func countBits2(num int) []int {
	var res []int
	res = make([]int, num+1)
	res[0] = 0
	for i := 0; i <= num/2+1; i++ {
		// 已知当前位置i的1数量为xi
		// 所有i+2^m位置的2数量为xi+1（注意2^m要大于i）
		prefix := 1
		for i+prefix <= num {
			if prefix > i {
				res[i+prefix] = res[i] + 1
			}
			prefix = prefix * 2
		}
	}
	return res
}

func main() {
	fmt.Println(countBits2(5))
}
