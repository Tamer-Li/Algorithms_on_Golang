package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in, out := inOut()
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscanln(in, &n)

		servers := parseServer(in, n)

		fmt.Fscanln(in, &m)

		images := parseImages(in, m)

		minDiff, serverAssignments := distributeImages(servers, images)
		fmt.Fprintln(out, minDiff)
		for _, server := range serverAssignments {
			fmt.Fprint(out, server, " ")
		}
		fmt.Fprintln(out)
	}
}

func inOut() (*bufio.Reader, *bufio.Writer) {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	return in, out
}

func parseServer(in *bufio.Reader, n int) []int {
	servers := make([]int, n)

	for i := 0; i < n; i++ {
		var s int
		fmt.Fscan(in, &s)
		servers[i] = s
	}

	return servers
}

func parseImages(in *bufio.Reader, m int) []int {
	images := make([]int, m)

	for i := 0; i < m; i++ {
		var s int
		fmt.Fscan(in, &s)
		images[i] = s
	}

	return images
}

func distributeImages(throughputs, weights []int) (int, []int) {
	sort.Ints(throughputs)
	sort.Ints(weights)

	m := len(weights)
	serverAssignments := make([]int, m)

	minTime := make([]int, m)
	maxTime := make([]int, m)

	for i := 0; i < m; i++ {
		minTime[i] = 1 << 30
		maxTime[i] = 0
	}

	for i := 0; i < m; i++ {
		weight := weights[i]
		bestServer := 0
		bestTime := 1 << 30

		for j := 0; j < len(throughputs); j++ {
			time := (weight + throughputs[j] - 1) / throughputs[j]
			if time < bestTime {
				bestTime = time
				bestServer = j + 1
			}
		}

		serverAssignments[i] = bestServer
		minTime[i] = bestTime
		maxTime[i] = bestTime
	}

	minDiff := 1 << 30
	for i := 0; i < m; i++ {
		if maxTime[i]-minTime[i] < minDiff {
			minDiff = maxTime[i] - minTime[i]
		}
	}

	return minDiff, serverAssignments
}
