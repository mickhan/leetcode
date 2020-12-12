package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	window := make(map[byte]int)
	var left, right, res int
	for right < len(s) {
		c := s[right]
		right++
		window[c]++
		for window[c] > 1 { // 大于1说明有重复元素，需要缩小窗口
			d := s[left]
			left++
			window[d]--
		}
		if right-left > res {
			res = right - left
		}
	}
	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}
