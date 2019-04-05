package main

import (
	"fmt"
)

func calcGroup(g string) int {
	switch g {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	case "IV":
		return 4
	case "IX":
		return 9
	case "XL":
		return 40
	case "XC":
		return 90
	case "CD":
		return 400
	case "CM":
		return 900
	}
	return 0
}

func romanToInt(s string) int {
	var sum int
	var x, y string
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			sum = sum + calcGroup(string(s[i]))
			continue
		}
		x, y = string(s[i]), string(s[i+1])
		if (x == "C" && (y == "D" || y == "M")) || (x == "X" && (y == "L" || y == "C")) || (x == "I" && (y == "V" || y == "X")) {
			sum = sum + calcGroup(string(x+y))
			i++
			continue
		}
		sum = sum + calcGroup(string(x))
	}
	return sum
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
