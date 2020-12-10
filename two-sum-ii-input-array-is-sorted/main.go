package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return []int{}
}

// 双指针
func twoSum2(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1 // 从数组两端开始
	for i != j {
		// numbers是升序的
		if numbers[i]+numbers[j] == target {
			return []int{i + 1, j + 1}
		} else if numbers[i]+numbers[j] > target {
			j-- // 如果结果过大，右指针向左移
		} else if numbers[i]+numbers[j] < target {
			i++ // 如果结果过小，左指针向右移
		}
	}
	return []int{}
}

func main() {
	fmt.Println(twoSum2([]int{2, 7, 11, 15}, 9))
}
