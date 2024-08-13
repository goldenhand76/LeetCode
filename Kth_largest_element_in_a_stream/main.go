package main

import (
	"sort"
)

type KthLargest struct {
	nums []int
	k    int
}

// Constructor initializes the object with the `k` largest elements from `nums`.
func Constructor(k int, nums []int) KthLargest {
	// Sort the nums slice to maintain the order.
	sort.Ints(nums)

	// If nums has more than k elements, keep only the largest k elements.
	if len(nums) > k {
		nums = nums[len(nums)-k:]
	}

	return KthLargest{
		nums: nums,
		k:    k,
	}
}

// Add adds a new value `val` to the stream and returns the k-th largest element.
func (this *KthLargest) Add(val int) int {
	// Use binary search to find the insertion index for `val`.
	idx := sort.Search(len(this.nums), func(i int) bool {
		return this.nums[i] >= val
	})

	// Insert `val` at the appropriate position in the sorted list.
	this.nums = append(this.nums, 0)
	copy(this.nums[idx+1:], this.nums[idx:])
	this.nums[idx] = val

	// If the list exceeds k elements, remove the smallest element.
	if len(this.nums) > this.k {
		this.nums = this.nums[1:]
	}

	// The k-th largest element is now the first element in the list.
	return this.nums[0]
}

func main() {
	// Example usage:
	k := 3
	nums := []int{4, 5, 8, 2}
	obj := Constructor(k, nums)
	println(obj.Add(3))  // returns 4
	println(obj.Add(5))  // returns 5
	println(obj.Add(10)) // returns 5
	println(obj.Add(9))  // returns 8
	println(obj.Add(4))  // returns 8
}
