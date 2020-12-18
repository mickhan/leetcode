package main

import "fmt"

// 合并成新数组，再找中位数
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var nums []int
	var i, j int
	for {
		if i >= len(nums1) {
			for j < len(nums2) {
				nums = append(nums, nums2[j])
				j++
			}
			break
		}
		if j >= len(nums2) {
			for i < len(nums1) {
				nums = append(nums, nums1[i])
				i++
			}
			break
		}
		if nums1[i] <= nums2[j] {
			nums = append(nums, nums1[i])
			i++
		} else {
			nums = append(nums, nums2[j])
			j++
		}
	}
	l := len(nums)
	if l%2 == 1 {
		return float64(nums[int(l/2)])
	}
	return float64(nums[l/2]+nums[l/2-1]) / 2
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 3}, []int{2}))
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
