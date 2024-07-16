package main

import "testing"

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		twoSum([]int{1, 2, 3, 4, 5}, 6)
	}
}
