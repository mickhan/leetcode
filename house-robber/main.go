package main

import "fmt"

// 动态规划，dp[i]代表抢劫0..i号房所能得到的最大值
// max(dp[i-2]+nums[i],dp[i-1]) -> dp[i]
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = nums[i]
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(nums)-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{2, 1, 1, 2}))
}
