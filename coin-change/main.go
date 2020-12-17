package main

import (
	"fmt"
	"math"
	"strings"
)

func arraySum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}

func arrayToString(arr []int) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(arr), " "), ":"), "[]")
}

// 回溯
// 直接回溯规模太大了，会超时/超内存
func backtrack(coins []int, remain int, res map[string][]int, cur []int) {
	if remain == 0 {
		// fmt.Println(cur)
		k := arrayToString(cur)
		if _, ok := res[k]; ok {
			return
		} else {
			tmp := make([]int, len(cur))
			copy(tmp, cur)
			res[k] = tmp
		}
	}
	if remain < 0 {
		return
	}
	for _, c := range coins {
		cur = append(cur, c)
		backtrack(coins, remain-c, res, cur)
		cur = cur[:len(cur)-1]
	}
}

func coinChange(coins []int, amount int) int {
	res := make(map[string][]int, 0)
	cur := make([]int, 0)
	backtrack(coins, amount, res, cur)
	minRes := math.MaxInt32
	for _, r := range res {
		if minRes > len(r) {
			minRes = len(r)
		}
	}
	if minRes == math.MaxInt32 {
		return -1
	}
	return minRes
}

// 递归
// 当待找的钱为remain的时候，最小结果是recursion(remain-c)+1
// 直接回溯规模太大了，会超时/超内存
// 使用subRes记录子remain的结果，如果再次遇到，可以直接使用
func recursion(coins []int, remain int, subRes map[int]int) int {
	if remain == 0 {
		return 0
	}
	if remain < 0 {
		return -1
	}
	if v, ok := subRes[remain]; ok {
		return v
	}
	minChange := math.MaxInt32
	for _, c := range coins {
		r := recursion(coins, remain-c, subRes)
		if r == -1 {
			continue
		}
		if r+1 < minChange {
			minChange = r + 1
		}
	}
	if minChange == math.MaxInt32 {
		subRes[remain] = -1
		return -1
	}
	subRes[remain] = minChange
	return minChange
}

func coinChange2(coins []int, amount int) int {
	subRes := make(map[int]int, 0)
	return recursion(coins, amount, subRes)
}

// 动态规划
func coinChange3(coins []int, amount int) int {
	dp := make(map[int]int)
	for i := 1; i <= amount; i++ {
		dp[i] = math.MaxInt32
		for _, c := range coins {
			if i-c < 0 {
				continue
			}
			dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-c])+1))
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func main() {
	fmt.Println(coinChange3([]int{1, 2, 5}, 11))
	fmt.Println(coinChange3([]int{186, 419, 83, 408}, 6249))
}
