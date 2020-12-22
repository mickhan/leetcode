package main

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	mem := make(map[string][]string)
	for _, str := range strs {
		// 每个str排序，生成唯一的key
		var tmp []string
		for _, s := range str {
			tmp = append(tmp, fmt.Sprint(s))
		}
		sort.Strings(tmp)
		key := strings.Join(tmp, "")
		mem[key] = append(mem[key], str)
	}
	var res [][]string
	for _, row := range mem {
		res = append(res, row)
	}
	return res
}

func main() {
	fmt.Println(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
}
