package main

import (
	"bufio"
	"fmt"
	"os"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var out *bufio.Writer
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	testCases(out)
}

func testCases(out *bufio.Writer) {
	nums := []int{1, 2, 3, 4, 5}
	listNodes := createListNode(nums)
	newHead := reverseList(listNodes)

	for newHead != nil {
		fmt.Fprintln(out, newHead.Val)
		newHead = newHead.Next
	}
}

func createListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{
		Val: nums[0],
	}
	current := head

	for i := 1; i < len(nums); i++ {
		current.Next = &ListNode{
			Val: nums[i],
		}
		current = current.Next
	}

	return head
}

//func reverseList(head *ListNode) *ListNode {
//	current := head.Next
//	head.Next = nil
//
//	for current != nil {
//		current :=
//			head.Next == current
//	}
//}
