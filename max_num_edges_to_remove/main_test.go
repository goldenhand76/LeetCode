package main

import "testing"

func BenchmarkMaxNumEdgesToRemove(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		maxNumEdgesToRemove(4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}})
	}
}
