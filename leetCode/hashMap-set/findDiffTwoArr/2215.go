package main

import "fmt"

func main() {
	test1 := [][]int{
		[]int{1, 2, 3},
		[]int{2, 4, 6},
	}

	test2 := [][]int{
		[]int{1, 2, 3, 3},
		[]int{1, 1, 2, 2},
	}

	fmt.Println(findDifference(test1[0], test1[1]))
	fmt.Println(findDifference(test2[0], test2[1]))
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	hashMap1 := make(map[int]struct{}, len(nums1))
	hashMap2 := make(map[int]struct{}, len(nums2))

	for _, v := range nums1 {
		hashMap1[v] = struct{}{}
	}

	for _, v := range nums2 {
		hashMap2[v] = struct{}{}
	}

	var res1 []int
	var res2 []int

	for v, _ := range hashMap1 {
		if _, ok := hashMap2[v]; !ok {
			res1 = append(res1, v)
		}
	}

	for v, _ := range hashMap2 {
		if _, ok := hashMap1[v]; !ok {
			res2 = append(res2, v)
		}
	}

	return [][]int{res1, res2}
}
