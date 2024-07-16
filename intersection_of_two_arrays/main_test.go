package main

import "testing"

func BenchmarkIntersect(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		intersect([]int{1, 2, 2, 1}, []int{2, 2})
	}
}
