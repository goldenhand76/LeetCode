package main

import "fmt"

func main() {
	numTeams([]int{2, 5, 3, 4, 1})
	numTeams([]int{2, 1, 3})
	numTeams([]int{1, 2, 3, 4})
}

func numTeams(rating []int) int {
	n := len(rating)
	if n < 3 {
		return 0
	}

	// Arrays to store the count of increasing and decreasing subsequences of length 2 ending at each index
	increasing := make([]int, n)
	decreasing := make([]int, n)

	count := 0

	// Fill the arrays
	for j := 1; j < n; j++ {
		for i := 0; i < j; i++ {
			if rating[i] < rating[j] {
				increasing[j]++
				// If we find an increasing subsequence ending at j, add the count of subsequences ending at i to the result
				count += increasing[i]
			} else if rating[i] > rating[j] {
				decreasing[j]++
				// If we find a decreasing subsequence ending at j, add the count of subsequences ending at i to the result
				count += decreasing[i]
			}
		}
	}
	fmt.Println(count)
	return count
}
