package main

import (
	"sort"
)

func frequencySort(nums []int) []int {
	items := make(map[int]int)

	// Count frequencies
	for _, v := range nums {
		items[v]++
	}

	// Get unique elements
	// unique := make([]int, 0, len(items))
	// for k := range items {
	// 	unique = append(unique, k)
	// }

	// Sort unique elements by frequency, then by value in descending order
	// sort.Slice(unique, func(i, j int) bool {
	// 	if items[unique[i]] == items[unique[j]] {
	// 		return unique[i] > unique[j]
	// 	}
	// 	return items[unique[i]] < items[unique[j]]
	// })

	sort.SliceStable(nums, func(i, j int) bool {
		if items[nums[i]] == items[nums[j]] {
			return nums[i] > nums[j]
		}
		return items[nums[i]] < items[nums[j]]
	})

	// Reconstruct the array
	// idx := 0
	// for _, v := range unique {
	// 	for k := 0; k < items[v]; k++ {
	// 		nums[idx] = v
	// 		idx++
	// 	}
	// }

	return nums
}

func main() {
	a := []int{2, 3, 1, 3, 2}
	frequencySort(a)

}
