package main

import (
	"fmt"
	"math"
)

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	// fmt.Println(m, n)
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = grid[i][j]
			// 0,0 不需计算
			if i == 0 && j == 0 {
				continue	
			}
			// 第一行 只能向右
			if i == 0 {
				res[i][j] = res[i][j-1]+res[i][j]
				continue
			}
			// 第一列 只能向下
			if j == 0 {
				res[i][j] = res[i-1][j]+res[i][j]
				continue
			}
			// 其他 min(向下,向右)
			res[i][j] = int(math.Min(float64(res[i-1][j]), float64(res[i][j-1]))) + res[i][j]
		}
	}
	// fmt.Printf("%+v", res)
	return res[m-1][n-1]
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum([][]int{{1, 2, 3}, {4, 5, 6}}))
}
