package main

import (
	"fmt"
	"strings"
)

var glen int

func genFull(fullSet *[]string, cur string) {
	if len(cur) == 2*glen {
		*fullSet = append(*fullSet, cur)
		return
	}
	option := "()"
	for i := 0; i < 2; i++ {
		cur = cur + string(option[i])
		genFull(fullSet, cur)
		cur = cur[:len(cur)-1]
	}
}

func validate(pa string) bool {
	pb := []byte(pa)
	if string(pb[0]) == ")" {
		return false
	}
	for {
		lenA := len(pa)
		pa = strings.ReplaceAll(pa, "()", "")
		lenB := len(pa)
		if lenA == lenB || lenB == 0 {
			break
		}
	}
	// 检查字符串是否为空
	// 为空说明都是成对的，有效
	if len(pa) == 0 {
		return true
	} else {
		return false
	}
}

func generateParenthesis(n int) []string {
	var fullSet []string
	var res []string
	var cur string
	glen = n
	// 生成所有可能的排列组合
	genFull(&fullSet, cur)
	// 判断其中有效的排列
	for _, v := range fullSet {
		if validate(v) {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}
