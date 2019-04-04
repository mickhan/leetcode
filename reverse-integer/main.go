package main

import "fmt"
import "math"

func reverse(x int) int {
	// 用一个数组保存每一位
	var seq []int
	var remain int
	remain = x % 10
	x = x / 10
	for {
		if remain == 0 && x == 0 {
			break
		}
		seq = append(seq, remain)
		remain = x % 10
		x = x / 10
	}
	// seq 已经是反转的了，拼接成完整的整数
	var ret int
	for index, value := range seq {
		ret = ret + value*int(math.Pow10(len(seq)-index-1))
	}

	if ret >= 2147483647 || ret <= -2147483648 {
		return 0
	} else {
		return ret
	}
}

func main() {
	fmt.Print(reverse(-2147483648))
}
