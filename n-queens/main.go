package main

import "fmt"
import "math"
import "strings"

type Node struct {
    x int
    y int
}
var res [][]Node 
var cur []Node
var gN int

func makeOptions(node Node, options []Node) (newOpt []Node) {
    // 根据node删除options中所有不符合的选项，构造新的选项集
    // 横 x相等的
    // 竖 y相等的
    // 斜线 x-y的绝对值相等
    // 已经尝试过的 x小于当前值的（因为options是按x递增排列）
    for _, n := range options {
        if n.x <= node.x || n.y == node.y || math.Abs(float64(n.x-node.x))==math.Abs(float64(n.y-node.y)) {
            continue
        }
        newOpt = append(newOpt, n)
    }
    return 
}

func formatRes() (ret [][]string) {
    // 将res中的node转换为字符串结果
    for _, r := range res {
        // 先放入二维数组
        d2 := make([][]string, gN)
        for i := range d2 {
            d2[i] = make([]string, gN)
        }
        for i := 0; i < gN; i++ {
            for j := 0; j < gN; j++ {
                d2[i][j] = "."
            }
        }
        for _, n := range r {
            d2[n.x][n.y] = "Q"
        }
        // 再转换成一维数组
        var d1 []string
        for _, row := range d2 {
            d1 = append(d1, strings.Join(row, ""))
        }
        ret = append(ret, d1)
    }
    return 
}

func backtrace(options []Node) {
    // fmt.Printf("%+v --- %+v\n", cur, options)
    if len(options) == 0 && len(cur) == gN {
        // fmt.Println("Got One!")
        tmp := make([]Node, gN)
        copy(tmp, cur) 
        res = append(res, tmp)
    }
    // 尝试所有options
    for _, node := range options {
        cur = append(cur, node)
        backtrace(makeOptions(node, options))
        cur = cur[:len(cur)-1]
    }
}

func solveNQueens(n int) [][]string {
    // 清理之前的结果
    res = nil 
    // 处理特殊值
    if n == 1 {
        return [][]string{{"Q"}}
    }
    if n <= 3 {
        return [][]string{}
    }
    gN = n
    var options []Node
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            options = append(options, Node{i, j})
        }
    }
    backtrace(options)
    // fmt.Printf("%+v\n", res)
    return formatRes()
}

func main() {
    fmt.Println(solveNQueens(5))
}
