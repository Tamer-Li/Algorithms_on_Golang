package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in, out := inOut()
	defer out.Flush()

	var num1, num2 int16

	fmt.Fscanln(in, &num1)
	fmt.Fscanln(in, &num2)

	fmt.Fprintln(out, num1+num2)
}

func inOut() (*bufio.Reader, *bufio.Writer) {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)

	return in, out
}
