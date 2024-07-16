package main

import "testing"

func BenchmarkMerge(b *testing.B) {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}
