package main

import (
	"fmt"
	"sort"
)

// 每个人是{h, k}
func reconstructQueue(people [][]int) [][]int {
	queue := make([][]int, 0)
	// people排序
	sort.Slice(people, func(i int, j int) bool {
		// 如果h相同，按k升序排列
		if people[i][0] == people[j][0] {
			return people[i][1] < people[j][1]
		}
		// 按h降序排列
		return people[i][0] > people[j][0]
	})
	// 每个人按k为下标，插入到结果中
	for _, p := range people {
		queue = append(queue, []int{})
		copy(queue[p[1]+1:], queue[p[1]:])
		queue[p[1]] = p
	}
	return queue
}

func main() {
	fmt.Println(reconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}))
}
