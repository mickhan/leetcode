package main

import "fmt"

// 中心扩展
// i,i为中心
// i,i+1为中心
func countSubstrings(s string) int {
	cnt := 0
	for i := 0; i < len(s); i++ {
		cnt += centerExpand(s, i, i)
		if i < len(s)-1 {
			cnt += centerExpand(s, i, i+1)
		}
	}
	return cnt
}

func centerExpand(s string, c1 int, c2 int) (cnt int) {
	for s[c1] == s[c2] {
		cnt++
		c1--
		c2++
		if c1 < 0 || c2 >= len(s) {
			return
		}
	}
	return
}

func main() {
	fmt.Println(countSubstrings("abc"))
	fmt.Println(countSubstrings("aaa"))
}
