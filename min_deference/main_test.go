package main

import "testing"

func BenchmarkMinDifference(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		minDifference([]int{1, 5, 0, 10, 14})
	}
}
