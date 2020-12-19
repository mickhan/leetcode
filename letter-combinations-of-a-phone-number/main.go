package main

import "fmt"

func backtrack(digits string, cur string, options map[string][]string, res *[]string) {
	if len(digits) == 0 {
		*res = append(*res, cur)
		return
	}
	d := digits[:1]
	for _, l := range options[d] {
		cur = cur + l
		backtrack(digits[1:], cur, options, res)
		cur = cur[:len(cur)-1]
	}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	d2l := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	res := make([]string, 0)
	backtrack(digits, "", d2l, &res)
	return res
}

func main() {
	fmt.Println(letterCombinations("23"))
}
