package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var in *bufio.Scanner
	var out *bufio.Writer

	in = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	t := scanInt(in)

	for i := 0; i < t; i++ {
		n := scanInt(in)
		scanYaml(in, n, out)
	}

}

func scanInt(in *bufio.Scanner) int {
	in.Scan()
	n, _ := strconv.Atoi(in.Text())
	return n
}

func scanYaml(in *bufio.Scanner, n int, out *bufio.Writer) {
	stack := make([]string, 0)
	prevLevel := 0
	needLine := false

	for i := 0; i < n; i++ {
		in.Scan()
		s := in.Text()
		level := (len(s) - len(strings.TrimSpace(s))) / 4
		stack = stack[:level]

		if s[len(s)-1] == ':' {
			stack = append(stack, strings.TrimSpace(s[:len(s)-1]))
		} else {
			if level != prevLevel {
				if needLine {
					fmt.Fprintln(out, "")
				}
				if len(stack) > 0 {
					fmt.Fprintf(out, "[%s]\n", strings.Join(stack, "."))
				}
			}
			needLine = true
			split := strings.Split(s, ":")
			fmt.Fprintf(out, "%s = %s\n", strings.TrimSpace(split[0]), strings.TrimSpace(split[1]))
		}
		prevLevel = level
	}
	fmt.Fprintln(out, "")
}
