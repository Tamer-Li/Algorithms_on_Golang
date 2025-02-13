package main

import "fmt"

func main() {
	c := Constructor()

	arr := []int{1, 2, 100, 3001, 3002}

	for _, v := range arr {
		fmt.Println(c.Ping(v))
	}
}

type RecentCounter struct {
	queue []int
	idx   int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: make([]int, 0),
	}
}

func (rc *RecentCounter) Ping(t int) int {
	rc.queue = append(rc.queue, t)

	for rc.queue[rc.idx] < t-3000 {
		rc.idx++
	}

	return len(rc.queue) - rc.idx
}
