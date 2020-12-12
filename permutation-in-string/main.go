package main

import (
	"fmt"
)

func checkInclusion(s1 string, s2 string) bool {
	need := make(map[byte]int)
	for _, s := range []byte(s1) {
		need[s]++
	}
	window := make(map[byte]int)
	var left, right, valid int

	for right < len(s2) {
		c := s2[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left == len(s1) { // 找s1的排列，中间不能有其他字符。所以长度等于s1。
			if valid == len(need) { // 长度相等，有效个数也相等，找到了
				return true
			}
			d := s2[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return false
}

func main() {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("ab", "eidboaoo"))
	fmt.Println(checkInclusion("abc", "ccccbbbbaaaa"))
	fmt.Println(checkInclusion("abcdxabcde", "abcdeabcdx"))
}
