package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateLinkList(t *testing.T) {
	l1 := createLinkList([]int{1, 2, 5})
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5}}}

	require.Equal(t, l1.Val, l2.Val)
	require.Equal(t, l1.Next.Val, l2.Next.Val)
	require.Equal(t, l1.Next.Next.Val, l2.Next.Next.Val)
}

func TestAddTwoNumber(t *testing.T) {
	l1 := createLinkList([]int{1, 2, 5})
	l2 := createLinkList([]int{1, 2, 5})
	l3 := addTwoNumbers(l1, l2)

	require.Equal(t, l3.Val, l1.Val+l2.Val)
	require.Equal(t, l3.Next.Val, l1.Next.Val+l2.Next.Val)
	require.Equal(t, l3.Next.Next.Val, 0)
	require.Equal(t, l3.Next.Next.Next.Val, 1)
}

func BenchmarkAddTwoNumber(b *testing.B) {
	l1 := createLinkList([]int{1, 2, 5})
	l2 := createLinkList([]int{1, 2, 5})

	for i := 0; i < b.N; i++ {
		addTwoNumbers(l1, l2)
	}
}
