package main

import "fmt"

func removeElement(nums []int, val int) int {
	var x, y int
	for y < len(nums) {
		if nums[y] != val {
			nums[x], nums[y] = nums[y], nums[x]
			x++
		}
		y++
	}
	return x
}

func main() {
	n := []int{0, 1, 2, 2, 3, 0, 4, 2}
	fmt.Println(removeElement(n, 2))
	fmt.Println(n)
}
