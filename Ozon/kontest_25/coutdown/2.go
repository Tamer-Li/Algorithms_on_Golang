package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, out := inOut()
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)

	for i := 0; i < t; i++ {
		var n string
		fmt.Fscanln(in, &n)
		answer := solution(n)
		fmt.Fprintln(out, answer)
	}
}

func inOut() (*bufio.Reader, *bufio.Writer) {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	return in, out
}
func solution(s string) int {
	if len(s) == 1 {
		return int(s[0]-'0') + 1
	}
	return (len(s)-1)*10 + int(s[0]-'0') - boolToInt(findLower(s))
}

func findLower(s string) bool {
	for i := 1; i < len(s); i++ {
		if s[i] > s[0] {
			return false
		}
		if s[i] < s[0] {
			return true
		}
	}
	return false
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
