package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func printLinkList(l *ListNode) {
	fmt.Println(l.Val)
	next := l.Next
	for next != nil {
		fmt.Println(next.Val)
		next = next.Next
	}
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

func main() {
	l1 := createLinkList([]int{2, 4, 3})
	l2 := createLinkList([]int{5, 6, 7})
	result := addTwoNumbers(l1, l2)
	printLinkList(result)
}

// addTwoNumbers adds two numbers represented by linked lists.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry int
	dummy := &ListNode{}
	current := dummy

	// Iterate through both linked lists.
	for l1 != nil || l2 != nil || carry != 0 {
		var x, y int
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		// Calculate the sum and the carry.
		sum := x + y + carry
		carry = sum / 10
		sum = sum % 10

		// Create a new node for the current digit of the result.
		current.Next = &ListNode{Val: sum}
		current = current.Next
	}

	return dummy.Next
}
