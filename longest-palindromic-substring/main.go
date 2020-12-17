package main

import "fmt"

// 动态规划
// dp[i][j] i至j是否是回文子串
// dp[i+1][j-1]+i/j位置的状态 -> dp[i][j]
func longestPalindrome(s string) string {
	var start, end int
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	// 单个字符一定是回文
	for i := 0; i < len(s); i++ {
		dp[i][i] = true
	}
	// 两个字符
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = true
			start, end = i, i+1
		}
	}
	// 多个字符，根据状态转移
	for j := 0; j < len(s); j++ {
		for i := 0; i < j-1; i++ {
			if s[i] != s[j] {
				dp[i][j] = false
			} else {
				dp[i][j] = dp[i+1][j-1]
				if dp[i][j] && end-start < j-i {
					start, end = i, j
				}
			}
		}
	}
	// for _, row := range dp {
	// 	fmt.Println(row)
	// }
	return s[start : end+1]
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}
