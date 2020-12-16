package main

import "fmt"

func fib(N int) int {
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 1
	}
	return fib(N-1) + fib(N-2)
}

func fib2(N int) int {
	a, b := 0, 1
	for i := 0; i < N; i++ {
		a, b = b, a+b
	}
	return a
}

func main() {
	fmt.Println(fib2(0))
	fmt.Println(fib2(1))
	fmt.Println(fib2(4))
}
