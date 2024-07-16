package main

import (
	"math/rand/v2"
	"testing"
)

func BenchmarkQuickFindUF(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		uf := NewQuickFindUF(100)
		for range 100 {
			uf.Union(rand.IntN(100), rand.IntN(100))
		}
		for range 100 {
			uf.connected(rand.IntN(100), rand.IntN(100))
		}
	}

}
