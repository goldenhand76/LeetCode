package main

import (
	"fmt"
	"sort"
)

func removeElement(nums []int, val int) int {
	k := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k

}

func main() {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val := 2
	expectedNums := []int{0, 1, 3, 0, 4} // Example valid expected result

	k := removeElement(nums, val)
	fmt.Println("k =", k)
	fmt.Println("First k elements of nums:", nums[:k])

	sort.Ints(nums[:k])
	sort.Ints(expectedNums)

	// Custom judge check
	if k != len(expectedNums) {
		panic("Length mismatch!")
	}
	for i := 0; i < k; i++ {
		if nums[i] != expectedNums[i] {
			panic(fmt.Sprintf("Mismatch at index %d: got %d, expected %d", i, nums[i], expectedNums[i]))
		}
	}

	fmt.Println("Test passed!")
}
