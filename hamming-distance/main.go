package main

import (
	"fmt"
	"strconv"
)

func formatInt(n int) string {
	return strconv.FormatInt(int64(n), 2)
}

func hammingDistance(x int, y int) int {
	xBin := formatInt(x)
	yBin := formatInt(y)
	var hamDist int
	if len(xBin) > len(yBin) {
		yBin = fmt.Sprintf("%0*s", len(xBin), yBin)
	} else {
		xBin = fmt.Sprintf("%0*s", len(yBin), xBin)
	}
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
