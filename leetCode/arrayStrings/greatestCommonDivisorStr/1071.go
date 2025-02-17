package main

import (
	"fmt"
)

func main() {
	str1 := [...]string{
		"ABCABC",
		"ABABAB",
		"LEET",
	}
	str2 := [...]string{
		"ABC",
		"ABAB",
		"CODE",
	}
	out := [...]string{
		"ABC",
		"AB",
		"",
	}

	for i := 0; i < 3; i++ {
		res := gcdOfStrings(str1[i], str2[i])
		if res == out[i] {
			fmt.Printf("YES %d\n", i)
		} else {
			fmt.Printf("NO %d\n", i)
		}
	}
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	l := findLen(len(str1), len(str2))
	return str1[:l]
}

func findLen(num1, num2 int) int {
	if num2 == 0 {
		return num1
	}

	return findLen(num2, num1%num2)
}
