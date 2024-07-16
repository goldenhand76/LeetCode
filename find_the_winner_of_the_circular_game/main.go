package main

import "fmt"

type LinkList struct {
	val      int
	next     *LinkList
	previous *LinkList
}

func main() {
	result := findTheWinner(5, 2)
	fmt.Print(result)
}

func createLinkedList(n int) *LinkList {
	root := &LinkList{val: 1}
	linkList := root
	for i := 2; i <= n; i++ {
		linkList.next = &LinkList{val: i, previous: linkList}
		linkList = linkList.next
	}
	linkList.next = root
	root.previous = linkList
	return root
}

func (linkedList *LinkList) removeNode() *LinkList {
	linkedList.previous.next = linkedList.next
	linkedList.next.previous = linkedList.previous
	return linkedList.next
}

func findTheWinner(n int, k int) int {
	l1 := createLinkedList(n)

	for l1.next != l1 {
		for i := 0; i < k-1; i++ {
			l1 = l1.next
		}
		l1 = l1.removeNode()
	}
	return l1.val
}
