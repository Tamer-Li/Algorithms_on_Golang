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
		var n int
		fmt.Fscanln(in, &n)

		arrStr := parseStr(in, n)
		answer := compareArr(arrStr, n)
		fmt.Fprintln(out, answer)
	}
}

func inOut() (*bufio.Reader, *bufio.Writer) {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	return in, out
}

func parseStr(in *bufio.Reader, n int) []string {
	arrStr := make([]string, n)

	for i := 0; i < n; i++ {
		var str string
		fmt.Fscanln(in, &str)
		arrStr[i] = str
	}

	return arrStr
}

func compareArr(arr []string, n int) int {
	count := 0
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if compareStr(arr[i], arr[j]) {
				count++
			}
		}
	}
	return count
}

func compareStr(str1, str2 string) bool {
	diff := len(str1) - len(str2)

	if diff < 0 {
		diff = -diff
	}

	if int(diff) > 1 {
		return false
	}

	if len(str1) == 1 || len(str2) == 1 {
		if len(str1) == 1 && len(str2) == 1 {
			return str1[0] == str2[0]
		}
		if len(str1) == 1 {
			return checkSingleChar(str1, str2)
		}
		if len(str2) == 1 {
			return checkSingleChar(str2, str1)
		}
	}

	if diff == 1 {
		return checkShort(str1, str2)
	}

	evenMatch := true
	for i := 0; i < len(str1) && i < len(str2); i += 2 {
		if str1[i] != str2[i] {
			evenMatch = false
			break
		}
	}

	oddMatch := true
	for i := 1; i < len(str1) && i < len(str2); i += 2 {
		if str1[i] != str2[i] {
			oddMatch = false
			break
		}
	}

	return evenMatch || oddMatch
}

func checkSingleChar(singleStr, longStr string) bool {
	if singleStr[0] != longStr[0] {
		return false
	}

	for i := 2; i < len(longStr); i += 2 {
		if longStr[i] != longStr[0] {
			return false
		}
	}

	return true
}

func checkShort(str1, str2 string) bool {
	var minLen int
	answer := true
	if len(str1) < len(str2) {
		minLen = len(str1)
	} else {
		minLen = len(str2)
	}

	if minLen%2 == 1 {
		for i := 0; i < len(str1) && i < len(str2); i += 2 {
			if str1[i] != str2[i] {
				answer = false
				break
			}
		}
	} else {
		for i := 1; i < len(str1) && i < len(str2); i += 2 {
			if str1[i] != str2[i] {
				answer = false
				break
			}
		}
	}

	return answer
}
