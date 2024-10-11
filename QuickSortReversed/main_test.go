package main

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSort(t *testing.T) {
	for i := 0; i <= 100; i++ {
		a := rand.Perm(100)
		e := 0
		expected := make([]int, len(a))
		copy(expected, a)

		sort.Sort(sort.Reverse(sort.IntSlice(expected)))

		ReversedQuickSort(&e, &a, 0, len(a)-1)

		require.Equal(t, expected, a)
	}

}

func benchmarkReversedQuickSort(b *testing.B, size int) {
	a := rand.Perm(size)
	totalExchanges := 0
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		tmp := make([]int, len(a))
		e := 0
		copy(tmp, a)
		ReversedQuickSort(&e, &a, 0, len(a)-1)
		totalExchanges += e
	}

	meanExchanges := totalExchanges / b.N
	b.ReportMetric(float64(meanExchanges), "exchange/op")
}

func BenchmarkReversedQuickSort_100(b *testing.B)    { benchmarkReversedQuickSort(b, 100) }
func BenchmarkReversedQuickSort_1000(b *testing.B)   { benchmarkReversedQuickSort(b, 1000) }
func BenchmarkReversedQuickSort_10000(b *testing.B)  { benchmarkReversedQuickSort(b, 10000) }
func BenchmarkReversedQuickSort_100000(b *testing.B) { benchmarkReversedQuickSort(b, 100000) }
