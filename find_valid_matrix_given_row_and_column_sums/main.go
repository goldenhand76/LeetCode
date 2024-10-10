package main

import "fmt"

func main() {
	ans := restoreMatrix([]int{1, 0}, []int{1})
	fmt.Println(ans)
}

func restoreMatrix(rowSum []int, colSum []int) [][]int {
	var answer = make([][]int, len(rowSum))
	for i := range answer {
		answer[i] = make([]int, len(colSum))
	}

	for size := range min(len(rowSum), len(colSum)) {
		for i := 0; i <= size; i++ {
			answer[i][size] = min(rowSum[i], colSum[size])
			colSum[size] -= answer[i][size]
			rowSum[i] -= answer[i][size]
		}

		for i := 0; i <= size; i++ {
			if i == size {
				continue
			}
			answer[size][i] = min(rowSum[size], colSum[i])
			rowSum[size] -= answer[size][i]
			colSum[i] -= answer[size][i]
		}
	}
	if len(rowSum) > len(colSum) {
		for size := len(colSum); size < len(rowSum); size++ {
			for i := range len(colSum) {
				answer[size][i] = min(rowSum[size], colSum[i])
				colSum[i] -= answer[size][i]
				rowSum[size] -= answer[size][i]
			}
		}
	} else if len(colSum) > len(rowSum) {
		for size := len(rowSum); size < len(colSum); size++ {
			for i := range len(rowSum) {
				answer[i][size] = min(rowSum[i], colSum[size])
				colSum[size] -= answer[i][size]
				rowSum[i] -= answer[i][size]
			}
		}
	}

	return answer
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
