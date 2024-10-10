package main

import "fmt"

func main() {
	minSwaps([]int{1, 1, 0, 0, 1})
}

func minSwaps(nums []int) int {
	n := len(nums)
	ones := make([]int, 2*n)
	count := 0
	for i := range n {
		if nums[i] == 1 {
			count++
		}
		ones[i] = count
	}
	new_count := count
	for i := n; i < 2*n; i++ {
		if nums[i-n] == 1 {
			new_count++
		}
		ones[i] = new_count
	}

	maxDiff := 0
	for i := range n {
		fmt.Println(ones[i : i+count])
		if ones[i+count] != ones[i] {
			maxDiff = max(maxDiff, ones[i+count]-ones[i])
		} else {
			if ones[i] == 1 {
				maxDiff = max(maxDiff, ones[i+count]-ones[i]-1)
			} else {
				maxDiff = max(maxDiff, ones[i+count]-ones[i]+1)
			}
		}

	}

	if count-maxDiff < 0 {
		return 0
	}
	return count - maxDiff
}
