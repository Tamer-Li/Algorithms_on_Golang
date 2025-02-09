package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	answer := testCase1()
	fmt.Fprintln(out, answer)

	answer = testCase2()
	fmt.Fprintln(out, answer)
}

func testCase1() float64 {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	return findMaxAverage(nums, k)
}

func testCase2() float64 {
	nums := []int{5}
	k := 1
	return findMaxAverage(nums, k)
}

func findMaxAverage(nums []int, k int) float64 {
	var left, right = 0, 0
	var sum float64
	var maxAvg = math.Inf(-1)

	for right < len(nums) {
		var window = right - left + 1
		sum += float64(nums[right])

		if window == k {
			maxAvg = max(sum/float64(k), maxAvg)
			sum -= float64(nums[left])
			left++
		}

		right++
	}
	return maxAvg
}
