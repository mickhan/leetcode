package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// 哈希
func majorityElement(nums []int) int {
	var major, max int
	cnt := make(map[int]int, 0)
	for _, i := range nums {
		// 记录每个数出现次数
		cnt[i]++
		// 维护最大值
		if cnt[i] >= max {
			max = cnt[i]
			major = i
		}
	}
	return major
}

// 排序
func majorityElement2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 随机
func majorityElement3(nums []int) int {
	// 生成随机位置序列
	pos := make([]int, len(nums))
	for i := range pos {
		pos[i] = i
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(pos), func(i, j int) { pos[i], pos[j] = pos[j], pos[i] })
	// 按序列依次尝试nums的各个位置
	// 看看该位置是否是多数元素
	for _, i := range pos {
		cnt := 0
		for _, n := range nums {
			if n == nums[i] {
				cnt++
			}
		}
		if float64(cnt) >= float64(len(nums))/2 {
			return nums[i]
		}
	}
	return 0
}

// 随机
func majorityElement4(nums []int) int {
	// 随机选取一个位置
	for {
		i := rand.Intn(len(nums))
		cnt := 0
		for _, n := range nums {
			if n == nums[i] {
				cnt++
			}
		}
		if float64(cnt) >= float64(len(nums))/2 {
			return nums[i]
		}
	}
}

func main() {
	fmt.Println(majorityElement4([]int{3, 3, 4}))
}
