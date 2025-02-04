package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	in, out := inOut()
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)

	for i := 0; i < t; i++ {
		var n int

		fmt.Fscanln(in, &n)
		answer := scanCargo(in, n)
		friend, err := parseList(in)
		if !err {
			fmt.Fprintln(out, "NO")
			continue
		}

		otvet := compareMaps(answer, friend)

		fmt.Fprintln(out, otvet)
	}

}

func inOut() (*bufio.Reader, *bufio.Writer) {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	return in, out
}

func scanCargo(in *bufio.Reader, n int) map[int][]string {
	list := make(map[int][]string)

	for i := 0; i < n; i++ {
		var price int
		var name string

		fmt.Fscanln(in, &name, &price)
		list[price] = append(list[price], name)
	}

	return list
}

func parseList(in *bufio.Reader) (map[int]string, bool) {
	var friend string
	fmt.Fscanln(in, &friend)

	if strings.ContainsAny(friend, " \t\n\r") {
		return nil, false
	}

	list := make(map[int]string)

	strFriend := strings.Split(friend, ",")

	for _, str := range strFriend {
		s := strings.Split(str, ":")
		if len(s) != 2 {
			return nil, false
		}
		name := s[0]

		if !isDigit(s[1]) {
			return nil, false
		}

		if !isValidPrice(s[1]) {
			return nil, false
		}

		price, _ := strconv.Atoi(s[1])

		_, ok := list[price]
		if ok {
			return nil, false
		}
		list[price] = name
	}

	return list, true
}

func isDigit(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func isValidPrice(priceStr string) bool {
	if len(priceStr) > 1 && priceStr[0] == '0' {
		return false
	}
	return true
}

func compareMaps(my map[int][]string, friend map[int]string) string {
	if len(my) != len(friend) {
		return "NO"
	}

	for key, arrStr := range my {
		var ok bool
		name, exists := friend[key]
		if !exists {
			return "NO"
		}

		for _, s := range arrStr {
			if s == name {
				ok = true
			}
		}

		if !ok {
			return "NO"
		}
	}

	return "YES"
}
