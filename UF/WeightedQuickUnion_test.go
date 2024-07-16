package main

import (
	"math/rand/v2"
	"testing"
)

func BenchmarkWeightedUF(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		uf := NewWeightedQuickUnionUF(1000)
		for range 1000 {
			uf.union(rand.IntN(100), rand.IntN(100))
		}
		for range 1000 {
			uf.connected(rand.IntN(100), rand.IntN(100))
		}
	}
}
