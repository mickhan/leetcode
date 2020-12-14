package main

import "fmt"

func moveZeroes(nums []int) {
	for i := 0; i < len(nums)-1; i++ {
		x, y := i, i+1
		if nums[x] != 0 {
			continue
		}
		for y < len(nums) {
			// fmt.Println(x, y, nums)
			if nums[y] != 0 {
				nums[x], nums[y] = nums[y], nums[x]
				x = y
			}
			y++
		}
	}
}

func moveZeroes2(nums []int) {
	x, y := 0, 1
	for x < len(nums)-1 && y < len(nums) {
		// fmt.Println(x, y, nums)
		if nums[x] != 0 {
			x++
			y = x + 1
			continue
		}
		if nums[y] != 0 {
			nums[x], nums[y] = nums[y], nums[x]
		} else {
			y++
		}
	}
}

func moveZeroes3(nums []int) {
	x, y, n := 0, 0, len(nums)
	for y < n {
		// fmt.Println(x, y, nums)
		if nums[y] != 0 {
			nums[x], nums[y] = nums[y], nums[x]
			x++
		}
		y++
	}
}

func main() {
	// n := []int{0, 1, 0, 3, 12}
	n := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}
	moveZeroes(n)
	fmt.Println(n)
}
