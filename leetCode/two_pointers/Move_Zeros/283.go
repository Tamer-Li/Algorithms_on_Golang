package main

import (
	"bufio"
	"fmt"
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

func testCase1() []int {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	return nums
}

func testCase2() []int {
	nums := []int{0}
	moveZeroes(nums)
	return nums
}

func moveZeroes(nums []int) {
	idxNoZero := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[idxNoZero] = nums[idxNoZero], nums[i]
			idxNoZero++
		}
	}
}
