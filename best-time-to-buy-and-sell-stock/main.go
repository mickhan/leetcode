package main

import (
	"fmt"
	"math"
)

// 双指针，计算最大差值
func maxProfit(prices []int) int {
	profit := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			curProfit := prices[j] - prices[i]
			if curProfit > profit {
				profit = curProfit
			}
		}
	}
	return profit
}

// 历史最低点
func maxProfit2(prices []int) int {
	profit, lowest := 0, math.MaxInt32
	for _, price := range prices {
		if price < lowest {
			lowest = price
			continue
		}
		curProfit := price - lowest
		if curProfit > profit {
			profit = curProfit
		}
	}
	return profit
}

func main() {
	fmt.Println(maxProfit2([]int{7, 1, 5, 3, 6, 4}))
}
