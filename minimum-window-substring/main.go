package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) string {
	need := make(map[byte]int)
	window := make(map[byte]int)
	for _, b := range []byte(t) {
		need[b]++
	}
	var left, right, valid, start int
	length := math.MaxInt32
	for right < len(s) {
		c := s[right]
		right++
		if _, ok := need[c]; ok {
			window[c]++
			if window[c] == need[c] {
				valid++
			}
		}
		for valid == len(need) {
			if right-left < length {
				start = left
				length = right - left
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
	if length < math.MaxInt32 {
		return s[start : start+length]
	}
	return ""
}

func main() {
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("a", "a"))
	fmt.Println(minWindow("abc", "z"))
}
