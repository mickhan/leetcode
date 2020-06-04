package main

import (
	"fmt"
	"math"
)

func minPathSum(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, m)
	}
	// 从右下角开始，逐渐填充结果矩阵
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			var down, right int
			res[i][j] = grid[i][j]
			// 向下
			if j == m-1 {
				down = math.MaxInt32
			} else {
				down = res[i][j+1]
			}
			// 向右
			if i == n-1 {
				right = math.MaxInt32
			} else {
				right = res[i+1][j]
			}
			// 选择其中较小的，加上本格的花费，作为本格的最优解
			if j != m-1 || i != n-1 {
				res[i][j] = int(math.Min(float64(down), float64(right))) + res[i][j]
			}
		}
	}
	return res[0][0]
}

func main() {
	fmt.Println(minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}))
	fmt.Println(minPathSum([][]int{{1, 2, 5}, {3, 2, 1}}))
}
