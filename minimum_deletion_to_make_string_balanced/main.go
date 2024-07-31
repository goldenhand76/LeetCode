package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minimumDeletions("b"))                                               // 0
	fmt.Println(minimumDeletions("aababbab"))                                        // 2
	fmt.Println(minimumDeletions("baababbaabbaaabaabbabbbabaaaaaabaabababaaababbb")) // 16

	fmt.Println(minimumDeletions_2("a"))                                               // 0
	fmt.Println(minimumDeletions_2("aababbab"))                                        // 2
	fmt.Println(minimumDeletions_2("baababbaabbaaabaabbabbbabaaaaaabaabababaaababbb")) // 16
}

func minimumDeletions(s string) int {
	// Calculate the prefix sums of 'b's
	prefixB := make([]int, len(s)+1)
	for i := 0; i < len(s); i++ {
		prefixB[i+1] = prefixB[i]
		if s[i] == 'b' {
			prefixB[i+1]++
		}
	}

	// Calculate the suffix sums of 'a's
	suffixA := make([]int, len(s)+1)
	for i := len(s) - 1; i >= 0; i-- {
		suffixA[i] = suffixA[i+1]
		if s[i] == 'a' {
			suffixA[i]++
		}
	}

	// Find the minimum deletions needed
	minDeletions := math.MaxInt32
	for i := 0; i <= len(s); i++ {
		minDeletions = min(minDeletions, prefixB[i]+suffixA[i])
	}

	return minDeletions
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumDeletions_2(s string) int {
	var res, bCount int

	for i := range s {
		if s[i] == 'a' {
			res = min(res+1, bCount)
		} else {
			bCount++
		}
	}
	return res
}
