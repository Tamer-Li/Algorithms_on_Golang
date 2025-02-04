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
		var n, m int
		fmt.Fscanln(in, &n, &m)
		answer := light(n, m)
		fmt.Fprintln(out, len(answer))
		for i := range answer {
			fmt.Fprintln(out, answer[i])
		}
	}
}

func inOut() (*bufio.Reader, *bufio.Writer) {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	return in, out
}

func light(n, m int) []string {
	x1, y1 := 1, 1
	var dir1 string
	var answer []string
	if n == 1 {
		dir1 = "R"
		answer = append(answer, fmt.Sprintf("%d %d %s", x1, y1, dir1))
		return answer
	}
	if m == 1 {
		dir1 = "D"
		answer = append(answer, fmt.Sprintf("%d %d %s", x1, y1, dir1))
		return answer
	}

	dir1 = "D"
	x2, y2 := 1, 2
	dir2 := "R"

	answer = append(answer, fmt.Sprintf("%d %d %s", x1, y1, dir1))
	answer = append(answer, fmt.Sprintf("%d %d %s", x2, y2, dir2))
	return answer
}
