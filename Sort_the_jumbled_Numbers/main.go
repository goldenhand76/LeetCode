package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	mapping := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	r := sortJumbled(mapping, nums)
	fmt.Println(r)

}

func sortJumbled(mapping []int, nums []int) []int {
	var findDigits func(int) int

	findDigits = func(num int) int {
		if num == 0 {
			return mapping[0]
		}
		transformedNum := 0
		multiplier := 1

		for num > 0 {
			digit := num % 10
			transformedNum += mapping[digit] * multiplier
			num /= 10
			multiplier *= 10
		}

		return transformedNum
	}

	sort.SliceStable(nums, func(i, j int) bool {
		return findDigits(nums[i]) < findDigits(nums[j])
	})

	return nums
}
