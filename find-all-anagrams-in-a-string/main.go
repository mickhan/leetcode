package main

import "fmt"

func findAnagrams(s string, p string) []int {
	var res []int
	need := make(map[byte]int)
	for _, b := range []byte(p) {
		need[b]++
	}
	window := make(map[byte]int)
	var left, right, valid int

	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for right-left == len(p) { // 找p的排列，中间不能有其他字符。所以长度等于p。
			if valid == len(need) { // 长度相等，有效个数也相等，找到了
				res = append(res, left)
			}
			d := s[left]
			left++
			if _, ok := need[d]; ok {
				if window[d] == need[d] {
					valid--
				}
				window[d]--
			}
		}
	}
	return res
}

func main() {
	fmt.Println(findAnagrams("cbaebabacd", "abc"))
}
