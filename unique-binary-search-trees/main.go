package main

import "fmt"

// 0个数 -> 1种树
// 1个数 -> 1种树
// 2个数 -> 第一个数为根，右边1个数有1种树；第二个数为根，左边1个数有1种树。共2种。
// 3个数 -> 第一个数为根，右边2个数有2种树；第二个数为根，左边1个数1种树，右边1个数1种树；第三个数为根，左边2个数有2种树。共5种树。
// 4个数 -> 第一个数为根，右边3个数有5种树；第二个树为根……
// 可以分解成子问题，用动态规划来解
func numTrees(n int) int {
    treeNum := make([]int, n+1)
    treeNum[0], treeNum[1] = 1, 1
    for i := 2; i <= n; i++ {
        for j := 1; j <= i; j++ {
            // 1...j...i，遍历1~n的每个位置
            // 每个位置的组合数量等于1...j的树数量乘j...i的树数量
            treeNum[i] += treeNum[j-1]*treeNum[i-j]
        }
    }
    return treeNum[n]
}

func main() {
    fmt.Println(numTrees(3))
}
