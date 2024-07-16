package main

import (
	"fmt"
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// l1 := createLinkList([]int{5, 3, 1, 2, 5, 1, 2})
	// result := nodesBetweenCriticalPoints(l1)
	// fmt.Println(result)
	// l2 := createLinkList([]int{1, 3, 2, 2, 3, 2, 2, 2, 7})
	// result = nodesBetweenCriticalPoints(l2)
	// fmt.Println(result)
	l3 := createLinkList([]int{2, 3, 3, 2})
	result := nodesBetweenCriticalPoints(l3)
	fmt.Println(result)
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	indexes := []int{}
	index := 2

	if head.Next.Next == nil {
		return []int{-1, -1}
	}

	for head.Next.Next != nil {
		if min(head.Val, head.Next.Val, head.Next.Next.Val) == head.Next.Val || max(head.Val, head.Next.Val, head.Next.Next.Val) == head.Next.Val {
			if head.Val != head.Next.Val && head.Next.Val != head.Next.Next.Val {
				indexes = append(indexes, index)
			}
		}
		index += 1
		head = head.Next
	}
	if len(indexes) < 2 {
		return []int{-1, -1}
	}
	sort.Ints(indexes)
	mindiff := 100000
	maxdiff := indexes[len(indexes)-1] - indexes[0]
	for i, v := range indexes {
		for _, v2 := range indexes[i+1:] {
			if v2-v <= mindiff {
				mindiff = v2 - v
			} else {
				break
			}
		}
	}
	return []int{mindiff, maxdiff}
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
