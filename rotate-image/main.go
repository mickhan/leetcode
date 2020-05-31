package main

import "fmt"

// 矩阵坐标映射关系
// x' = y
// y' = n - 1 - x
func rotate1(matrix [][]int) {
	n := len(matrix[0])
	m2 := make([][]int, n) // 不可以使用另一个矩阵来旋转图像，这个方法只是提供思路
	for i := range m2 {
		m2[i] = make([]int, n)
	}
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			xn := y
			yn := n - 1 - x
			m2[xn][yn] = matrix[x][y]
		}
	}
	fmt.Println(m2)
}

// 划分成4个小矩阵，映射到自身
// 一轮操作映射4次
func rotate(matrix [][]int) {
	n := len(matrix[0])
	// 矩阵划分：偶数刚好等分，奇数x多分1
	for x := 0; x < n/2+n%2; x++ {
		for y := 0; y < n/2; y++ {
			tmp := matrix[n-1-y][x]
			matrix[n-1-y][x] = matrix[n-1-x][n-1-y]
			matrix[n-1-x][n-1-y] = matrix[y][n-1-x]
			matrix[y][n-1-x] = matrix[x][y]
			matrix[x][y] = tmp
		}
	}
	fmt.Println(matrix)
}

func main() {
	m := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	rotate(m)
}
