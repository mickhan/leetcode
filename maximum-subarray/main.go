package main

import (
	"fmt"
	"math"
)

// 动态规划，sum[i][j]代表nums[i...j]的序列和
// sum[i+1][j-1]+nums[i]+nums[j]->sum[i][j]
// O(n^2)会超时/超内存
func maxSubArray(nums []int) int {
	max := math.MinInt32
	sum := make([][]int, len(nums))
	for s := range sum {
		sum[s] = make([]int, len(nums))
	}
	for i := 0; i < len(nums); i++ {
		sum[i][i] = nums[i]
		if sum[i][i] > max {
			max = sum[i][i]
		}
	}
	for i := 0; i < len(nums)-1; i++ {
		sum[i][i+1] = nums[i] + nums[i+1]
		if sum[i][i+1] > max {
			max = sum[i][i+1]
		}
	}
	for j := 0; j < len(nums); j++ {
		for i := 0; i < j-1; i++ {
			sum[i][j] = nums[i] + nums[j] + sum[i+1][j-1]
			if sum[i][j] > max {
				max = sum[i][j]
			}
		}
	}
	return max
}

// 动态规划，dp[i]代表nums[x...i]的序列和
// dp[i-1] nums[i] -> dp[i]
func maxSubArray2(nums []int) int {
	res := nums[0]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		if dp[i] > res {
			res = dp[i]
		}
	}
	return res
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxSubArray2([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
