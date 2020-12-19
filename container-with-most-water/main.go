package main

import "fmt"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 遍历所有可能性
func maxArea(height []int) int {
	var res int
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			tmp := min(height[i], height[j]) * (j - i)
			if tmp > res {
				res = tmp
			}
		}
	}
	return res
}

// 双指针
func maxArea2(height []int) int {
	l, r, res := 0, len(height)-1, 0
	for l < r {
		// 此时r-l已经是最长距离
		// min(height[l], height[r])已经得到了较低的一边
		// 如果此时移动较高的一边，计算出的容积不可能更大，只可能更小
		tmp := min(height[l], height[r]) * (r - l)
		if tmp > res {
			res = tmp
		}
		// 移动较低的一边
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}

func main() {
	fmt.Println(maxArea2([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}
