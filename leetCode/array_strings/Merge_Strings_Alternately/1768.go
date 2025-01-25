package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	answer := testCase1()
	fmt.Fprintln(out, answer)

	answer = testCase2()
	fmt.Fprintln(out, answer)

	answer = testCase3()
	fmt.Fprintln(out, answer)
}

func testCase1() string {
	word1 := "abc"
	word2 := "pqr"
	return mergeAlternately(word1, word2)
}

func testCase2() string {
	word1 := "ab"
	word2 := "pqrs"
	return mergeAlternately(word1, word2)
}

func testCase3() string {
	word1 := "abcd"
	word2 := "pq"
	return mergeAlternately(word1, word2)
}

func mergeAlternately(word1 string, word2 string) string {
	var builder strings.Builder

	l1 := len(word1)
	l2 := len(word2)
	minLen := l1
	if l2 < l1 {
		minLen = l2
	}

	for i := 0; i < minLen; i++ {
		builder.WriteByte(word1[i])
		builder.WriteByte(word2[i])
	}

	if l1 > l2 {
		builder.WriteString(word1[minLen:])
	} else if l2 > l1 {
		builder.WriteString(word2[minLen:])
	}

	return builder.String()
}

func mergeAlternatelyOfBest(word1 string, word2 string) string {
	runes1 := []rune(word1)
	runes2 := []rune(word2)
	n, m := len(runes1), len(runes2)
	result := make([]rune, 0, n+m)

	for i := 0; i < max(n, m); i++ {
		if i < n {
			result = append(result, runes1[i])
		}
		if i < m {
			result = append(result, runes2[i])
		}
	}
	return string(result)
}
