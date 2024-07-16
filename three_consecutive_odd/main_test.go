package main

import (
	"testing"
)

func BenchmarkThreeConsecutiveOdds(b *testing.B) {
	for i := 0; i < b.N; i++ {
		threeConsecutiveOdds([]int{1, 2, 34, 3, 4, 5, 7, 23, 12})
	}
}
