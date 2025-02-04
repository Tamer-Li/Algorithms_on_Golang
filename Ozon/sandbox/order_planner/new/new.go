package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Truck struct {
	index    int
	start    int
	end      int
	capacity int
}

type Departure struct {
	time  int
	index int
	truck int
}

func main() {
	in, out := inOut()
	defer out.Flush()

	var t int
	fmt.Fscanln(in, &t)

	for i := 0; i < t; i++ {
		var n, m int
		fmt.Fscanln(in, &n)
		arrival := scanDeparture(in, n)
		fmt.Fscanln(in, &m)
		trucks := scanTruck(in, m)
	}
}

func inOut() (*bufio.Reader, *bufio.Writer) {
	var in *bufio.Reader
	var out *bufio.Writer

	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)

	return in, out
}

func scanDeparture(in *bufio.Reader, n int) []Departure {
	arrival := make([]Departure, n)

	for i := 0; i < n; i++ {
		var time int
		fmt.Fscan(in, &time)
		arrival[i] = Departure{
			time:  time,
			index: i,
		}
	}

	sort.Slice(arrival, func(i, j int) bool {
		return arrival[i].time < arrival[j].time
	})

	return arrival
}

func scanTruck(in *bufio.Reader, m int) []Truck {
	trucks := make([]Truck, m)

	for j := 0; j < m; j++ {
		var start, end, capacity int
		fmt.Fscan(in, &start, &end, &capacity)
		trucks[j] = Truck{
			index:    j,
			start:    start,
			end:      end,
			capacity: capacity,
		}
	}

	sort.Slice(trucks, func(i, j int) bool {
		return trucks[i].start < trucks[j].start
	})

	return trucks
}

func loading(arrival []Departure, trucks []Truck) []Departure {
	for i := range arrival {
		arrival[i].truck = findTruck(arrival[i].time, &trucks)
	}
	return arrival
}

func findTruck(time int, trucks *[]Truck) int {
	for 

	return -1
}
