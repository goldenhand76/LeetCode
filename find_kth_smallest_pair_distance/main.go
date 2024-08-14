package main

import (
	"fmt"
	"sort"
)

// Helper function to count pairs with distance <= mid
func countPairs(nums []int, mid int) int {
	count := 0
	j := 0
	for i := 0; i < len(nums); i++ {
		for j < len(nums) && nums[j]-nums[i] <= mid {
			j++
		}
		count += j - i - 1
	}
	return count
}

func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)

	left, right := 0, nums[len(nums)-1]-nums[0]

	for left < right {
		mid := (left + right) / 2
		if countPairs(nums, mid) >= k {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	nums := []int{1, 3, 1}
	k := 1
	fmt.Println(smallestDistancePair(nums, k)) // Output: 0
}
