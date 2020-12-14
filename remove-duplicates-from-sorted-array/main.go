package main

import "fmt"

func removeDuplicates(nums []int) int {
	var x, y int
	for y < len(nums) {
		// fmt.Println(x, y, nums)
		if nums[x] != nums[y] {
			x++
			nums[x], nums[y] = nums[y], nums[x]
		}
		y++
	}
	return x + 1
}

func main() {
	n := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(n))
	fmt.Println(n)
}
