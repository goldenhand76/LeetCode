package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := createLinkList([]int{0, 3, 1, 0, 4, 5, 2, 0})
	mergeNodes(l1)
	l2 := createLinkList([]int{0, 1, 0, 3, 0, 2, 2, 0})
	mergeNodes(l2)
}

func mergeNodes(head *ListNode) *ListNode {
	sum := 0
	dummy := &ListNode{}
	output := dummy
	current := head
	for current.Next != nil {
		for current.Next.Val != 0 {
			if current.Val == 0 {
				sum = 0
			} else {
				sum += current.Val
			}
			current = current.Next
		}
		sum += current.Val
		output.Next = &ListNode{Val: sum}
		output = output.Next
		current = current.Next
	}

	fmt.Print(dummy.Next)
	return dummy.Next

}

func createLinkList(nums []int) *ListNode {
	head := &ListNode{Val: nums[0]}
	current := head
	for _, num := range nums[1:] {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head
}
