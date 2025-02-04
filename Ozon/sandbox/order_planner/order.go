package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Truck struct {
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
	var in *bufio.Scanner
	var out *bufio.Writer

	in = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	t := scanInt(in)

	for i := 0; i < t; i++ {
		n := scanInt(in)
		arrival := scanArr(in, n)
		m := scanInt(in)
		trucks := scanTruck(in, m)
		data := loading(arrival, trucks)
		output(out, data)
	}
}

func scanInt(in *bufio.Scanner) int {
	in.Scan()
	num, _ := strconv.Atoi(in.Text())
	return num
}

func scanArr(in *bufio.Scanner, n int) []Departure {
	arrival := make([]Departure, n)

	in.Scan()
	strArr := strings.Fields(in.Text())

	if len(strArr) != n {
		return nil
	}

	for i := 0; i < n; i++ {
		arrival[i].time, _ = strconv.Atoi(strArr[i])
		arrival[i].index = i
	}

	sort.Slice(arrival, func(i, j int) bool {
		return arrival[i].time < arrival[j].time
	})

	return arrival
}

func scanTruck(in *bufio.Scanner, m int) []Truck {
	arrTruck := make([]Truck, m)

	for j := 0; j < m; j++ {
		in.Scan()
		truck := strings.Fields(in.Text())

		if len(truck) != 3 {
			return nil
		}

		arrTruck[j].start, _ = strconv.Atoi(truck[0])
		arrTruck[j].end, _ = strconv.Atoi(truck[1])
		arrTruck[j].capacity, _ = strconv.Atoi(truck[2])
	}

	return arrTruck
}

func loading(arrival []Departure, trucks []Truck) []Departure {

	for i, d := range arrival {
		arrival[i].truck = interval(trucks, d.time)
	}

	return arrival
}

func interval(trucks []Truck, time int) int {
	for i := range trucks {
		if trucks[i].capacity > 0 && trucks[i].start <= time && time <= trucks[i].end {
			trucks[i].capacity--
			return i + 1
		}
	}
	return -1
}

func output(out *bufio.Writer, d []Departure) {
	sort.Slice(d, func(i, j int) bool {
		return d[i].index < d[j].index
	})
	strData := make([]string, len(d))
	for i, v := range d {
		strData[i] = strconv.Itoa(v.truck)
	}
	fmt.Fprintln(out, strings.Join(strData, " "))
}
