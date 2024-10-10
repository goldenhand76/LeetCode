package main

import (
	"sort"
)

func main() {
	combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 5)
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates) // Sort the candidates to handle duplicates
	var result [][]int
	var combination []int
	findCombinations(candidates, target, 0, combination, &result)
	return result
}

func findCombinations(candidates []int, target int, start int, combination []int, result *[][]int) {
	if target == 0 {
		// Make a copy of the combination and add it to the result
		temp := make([]int, len(combination))
		copy(temp, combination)
		*result = append(*result, temp)
		return
	}

	for i := start; i < len(candidates); i++ {
		// Skip duplicates
		if i > start && candidates[i] == candidates[i-1] {
			continue
		}

		// If the current candidate exceeds the target, break out of the loop
		if candidates[i] > target {
			break
		}

		// Include the current candidate in the combination
		combination = append(combination, candidates[i])

		// Recurse with the reduced target and next starting index
		findCombinations(candidates, target-candidates[i], i+1, combination, result)

		// Backtrack by removing the last element added
		combination = combination[:len(combination)-1]
	}
}
