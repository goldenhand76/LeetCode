package main

import "math"

func main() {
	minHeightShelves([][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}, 4)
}

func minHeightShelves(books [][]int, shelfWidth int) int {
	n := len(books)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0

	for i := 1; i <= n; i++ {
		current_width := 0
		max_height := 0
		for j := i; j > 0; j-- {
			current_width += books[j-1][0]
			if current_width > shelfWidth {
				break
			}
			max_height = max(max_height, books[j-1][1])
			dp[i] = min(dp[i], dp[j-1]+max_height)
		}
	}
	return dp[n]
}
