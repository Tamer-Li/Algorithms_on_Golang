package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var in *bufio.Scanner
	var out *bufio.Writer

	in = bufio.NewScanner(os.Stdin)
	buf := make([]byte, 0, 64*1024)
	in.Buffer(buf, 10*1024*1024)

	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	t := scanInt(in)

	for i := 0; i < t; i++ {
		var answer bool
		length := scanInt(in)

		arrNum := scanArrNum(in, length)

		arrCheck := scanArrStr(in, length)
		if arrCheck != "" {
			answer = check(arrNum, arrCheck)
		}

		output(out, answer)
	}
}

func scanInt(scan *bufio.Scanner) int {
	scan.Scan()
	num, _ := strconv.Atoi(scan.Text())
	return num
}

func scanArrNum(scan *bufio.Scanner, length int) []int {
	scan.Scan()
	str := scan.Text()

	arrStr := strings.Fields(str)
	arrNum := make([]int, 0, length)

	for _, s := range arrStr {
		num, _ := strconv.Atoi(s)
		arrNum = append(arrNum, num)
	}

	sort.Ints(arrNum)

	return arrNum
}

func scanArrStr(scan *bufio.Scanner, length int) string {
	scan.Scan()
	str := scan.Text()

	arrCheck := strings.Fields(str)
	if len(arrCheck) != length {
		return ""
	}

	return str
}

func check(arrNum []int, arrCheck string) bool {
	var arrStr strings.Builder
	for i, num := range arrNum {
		if i > 0 {
			arrStr.WriteString(" ")
		}
		arrStr.WriteString(strconv.Itoa(num))
	}

	return arrStr.String() == arrCheck
}

func output(out *bufio.Writer, answer bool) {
	if answer {
		fmt.Fprintln(out, "yes")
	} else {
		fmt.Fprintln(out, "no")
	}
}
