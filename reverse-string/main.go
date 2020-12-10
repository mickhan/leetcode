package main

import "fmt"

func reverseString(s []byte) {
	i, j := 0, len(s)-1
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func main() {
	s := []byte("hello")
	fmt.Println(s)
	reverseString(s)
	fmt.Println(s)
}
