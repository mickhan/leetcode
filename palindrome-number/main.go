package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var x_origin = x
	var remain int
	var x_rev int
	for {
		remain = x % 10
		x = x / 10
		if remain == 0 && x == 0 {
			break
		}
		x_rev = x_rev*10 + remain
	}
	// fmt.Printf("%v", pos)
	// fmt.Println(x, x_rev)
	if x_origin == x_rev {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(isPalindrome(121))
}
