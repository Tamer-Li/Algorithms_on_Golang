package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var s string

		fmt.Fscan(in, &s)

		if len(s) == 1 {
			fmt.Fprintln(out, 0)
			continue
		}

		removeDigit := -1

		for j := 0; j < len(s)-1; j++ {
			if s[j] < s[j+1] {
				removeDigit = j
				break
			}
		}

		if removeDigit == -1 {
			removeDigit = len(s) - 1
		}

		fmt.Fprint(out, s[:removeDigit])
		fmt.Fprintln(out, s[removeDigit+1:])
	}
}
