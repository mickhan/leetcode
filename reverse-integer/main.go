package main

import "fmt"

func reverse(x int) int {
	var remain int
	var ret int
	for {
		remain = x % 10
		x = x / 10
		if remain == 0 && x == 0 {
			break
		}
		ret = ret*10 + remain
	}
	if ret >= 2147483647 || ret <= -2147483648 {
		return 0
	} else {
		return ret
	}
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-12345))
	fmt.Println(reverse(-2147483648))
}
