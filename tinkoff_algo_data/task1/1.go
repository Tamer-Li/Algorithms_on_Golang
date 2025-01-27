package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, out := inOut()
	defer out.Flush()

	var n, a int

	fmt.Fscanln(in, &n, &a)

	friends := make([]int, n)

	for i := 0; i < n; i++ {
		var t int
		fmt.Fscan(in, &t)
		friends[i] = t
	}

	timeArr := timeDialog(friends, n, a)
	output(out, timeArr)
}

func inOut() (*bufio.Reader, *bufio.Writer) {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)

	return in, out
}

func timeDialog(friends []int, n, a int) []int {
	timeArr := make([]int, n)
	timeArr[0] = friends[0] + a

	if n > 1 {
		j := 0
		for _, t := range friends[1:] {
			timeMeet := timeArr[j]
			if timeMeet < t {
				timeMeet = t
			}
			j++
			timeArr[j] = timeMeet + a
		}
	}

	return timeArr
}

func output(out *bufio.Writer, arr []int) {
	for _, t := range arr {
		fmt.Fprintln(out, t)
	}
}
