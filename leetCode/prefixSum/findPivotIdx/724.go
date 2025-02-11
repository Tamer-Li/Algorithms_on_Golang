package main

import "fmt"

func main() {
	nums1 := []int{1, 7, 3, 6, 5, 6}
	nums2 := []int{1, 2, 3}
	nums3 := []int{2, 1, -1}

	fmt.Println(pivotIndex(nums1))
	fmt.Println(pivotIndex(nums2))
	fmt.Println(pivotIndex(nums3))
}

func pivotIndex(nums []int) int {
	sumRight := 0
	for _, num := range nums {
		sumRight += num
	}

	sumLeft := 0

	for i := 0; i < len(nums); i++ {
		sumRight -= nums[i]
		if sumLeft == sumRight {
			return i
		}
		sumLeft += nums[i]
	}

	return -1
}
