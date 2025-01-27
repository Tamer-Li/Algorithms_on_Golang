package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, out := inOut()
	defer out.Flush()

	var s string

	fmt.Fscanln(in, &s)

	arrSubStr := findSubstrings(s)
	score := scoreLetters(arrSubStr)
	output(out, score)
}

func inOut() (*bufio.Reader, *bufio.Writer) {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)

	return in, out
}

func findSubstrings(str string) []string {
	arrRune := []rune(str)
	l := len(arrRune)

	subStringArr := make([]string, 0)

	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			subStringArr = append(subStringArr, string(arrRune[i:j]))
		}
	}

	return subStringArr
}

func scoreLetters(arrSubStr []string) map[rune]int {
	mapScoreLet := make(map[rune]int)

	for _, subStr := range arrSubStr {
		for _, let := range subStr {
			if _, ok := mapScoreLet[let]; !ok {
				mapScoreLet[let] = 0
			}
			mapScoreLet[let]++
		}
	}

	return mapScoreLet
}

func output(out *bufio.Writer, m map[rune]int) {
	for letter, score := range m {
		fmt.Fprintf(out, "%c: %d\n", letter, score)
	}
}
