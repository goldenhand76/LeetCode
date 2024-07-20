package main

import (
	"fmt"
	"math/rand/v2"
	"testing"
)

func TestLuckyNumbers(t *testing.T) {
	col := rand.IntN(50)
	row := rand.IntN(50)
	var matrix [][]int = make([][]int, row)
	for i := range row {
		for _ = range col {
			matrix[i] = append(matrix[i], rand.IntN(10000))
		}
	}
	fmt.Print(matrix)
}
