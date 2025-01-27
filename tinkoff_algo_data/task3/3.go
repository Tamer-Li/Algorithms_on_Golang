package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Plant struct {
	height int
	index  int
}

func main() {
	in, out := inOut()
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)

	for i := 0; i < t; i++ {
		var n int
		fmt.Fscanln(in, &n)

		plantsHeight := parsePlants(in, n)
		plantsGrew := parsePlants(in, n)
		arrOleg := parsePlants(in, n)

		if !isPossible(n, plantsHeight, plantsGrew, arrOleg) {
			fmt.Fprintln(out, -1)
			continue
		}

		days := findMinDay(n, plantsHeight, plantsGrew, arrOleg)
		fmt.Fprintln(out, days)
	}
}

func inOut() (*bufio.Reader, *bufio.Writer) {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	return in, out
}

func parsePlants(in *bufio.Reader, n int) []int {
	arrPlants := make([]int, n)

	var str string
	fmt.Fscanln(in, &str)

	arrStr := strings.Fields(str)
	for i := 0; i < n; i++ {
		arrPlants[i], _ = strconv.Atoi(arrStr[i])
	}

	return arrPlants
}

func isPossible(n int, height, grew, oleg []int) bool {
	allSameGrowth := true
	for i := 1; i < n; i++ {
		if grew[i] != grew[0] {
			allSameGrowth = false
			break
		}
	}

	if allSameGrowth {
		return checkCondition(n, height, oleg)
	}

	return true
}

func checkCondition(n int, height, oleg []int) bool {
	plants := make([]Plant, n)
	for i := 0; i < n; i++ {
		plants[i] = Plant{height[i], i}
	}

	sort.Slice(plants, func(i, j int) bool {
		return plants[i].height < plants[j].height
	})

	for i := 0; i < n; i++ {
		if oleg[plants[i].index] != n-1-i {
			return false
		}
	}

	return true
}

func findMinDay(n int, height, grew, oleg []int) int {
	left, right := 0, 1000000
	result := -1

	for left <= right {
		mid := (left + right) / 2
		if isConditionMet(n, height, grew, oleg, mid) {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return result
}

func isConditionMet(n int, height, grew, oleg []int, days int) bool {
	newHeight := make([]int, n)
	for i := 0; i < n; i++ {
		newHeight[i] = height[i] + grew[i]*days
	}

	return checkCondition(n, newHeight, oleg)
}
