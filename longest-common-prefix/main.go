package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	var ret = ""
	var index = 0
	var index_char = ""
	if len(strs) == 0 {
		return ret
	}
	// 无限循环，直到出现不相等或者超过长度
	for {
		if index >= len(strs[0]) {
			break
		}
		index_char = string(strs[0][index])
		for _, s := range strs {
			if index >= len(s) || string(s[index]) != index_char {
				return ret
			}
		}
		index++
		ret = ret + index_char
	}
	return ret
}

func main() {
	strs := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}
