package main

import "fmt"

// 双重循环
func dailyTemperatures(T []int) []int {
	res := make([]int, len(T))
	for i := 0; i < len(T); i++ {
		for j := i + 1; j < len(T); j++ {
			if T[j] > T[i] {
				res[i] = j - i
				break
			}
		}
	}
	return res
}

// 单调栈
// 当T[i]小于栈顶的时候，将下标i入栈
// 当T[i]大于栈顶的时候，将所有小于T[i]的元素出栈，他们到i的距离就是res[x]。再将T[i]入栈
func dailyTemperatures2(T []int) []int {
	lt := len(T)
	stack := make([]int, 0)
	res := make([]int, lt)
	for i := 0; i < lt; i++ {
		iTop := len(stack) - 1
		for iTop >= 0 && T[stack[iTop]] < T[i] {
			res[stack[iTop]] = i - stack[iTop]
			stack = stack[:iTop]
			iTop = len(stack) - 1
		}
		stack = append(stack, i)
	}
	return res
}

func main() {
	fmt.Println(dailyTemperatures2([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
