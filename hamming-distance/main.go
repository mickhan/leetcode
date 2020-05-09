package main

import (
	"fmt"
	"strconv"
)

func formatInt(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func hammingDistance(x int, y int) int {
	// 把数字转换成二进制的字符串形式
	xBin := formatInt(x)
	yBin := formatInt(y)
	// 把较短的字符串用0填充，使长度相等
	if len(xBin) > len(yBin) {
		yBin = fmt.Sprintf("%0*s", len(xBin), yBin)
	} else {
		xBin = fmt.Sprintf("%0*s", len(yBin), xBin)
	}
	// 比较两个字符串的差异
	var hamDist int
	for i := 0; i < len(xBin); i++ {
		if xBin[i] != yBin[i] {
			hamDist++
		}
	}
	return hamDist
}

func main() {
	fmt.Println(hammingDistance(1, 4))
}
