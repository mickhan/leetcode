package main

import "fmt"

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	board := make([][]int, m)
	for i := range board {
		board[i] = make([]int, n)
	}
	board[0][1], board[1][0] = 1, 1
	// 第一行，只有一种走法
	for i := 0; i < n; i++ {
		board[0][i] = 1
	}
	// 第一列，只有一种走法
	for i := 0; i < m; i++ {
		board[i][0] = 1
	}
	// 其他，走法=从上面+从左边
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			board[i][j] = board[i-1][j] + board[i][j-1]
		}
	}
	fmt.Println(board)
	return board[m-1][n-1]
}

func main() {
	// fmt.Println(uniquePaths(3, 7))
	// fmt.Println(uniquePaths(1, 1))
	fmt.Println(uniquePaths(2, 2))
}
