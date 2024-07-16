package main

import "testing"

func BenchmarkMergeNodes(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		createLinkList([]int{0, 3, 1, 0, 4, 5, 2, 0})
	}
}
