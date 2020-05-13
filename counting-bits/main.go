package main

import "fmt"

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

func main() {
	fmt.Println(countBits(5))
}
