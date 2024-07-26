package main

import "fmt"

func main() {
	fmt.Println(sortArray([]int{5, 1, 1, 2, 0, 0}))
}

func sortArray(nums []int) []int {
	var sink func(int, int)
	N := len(nums)

	sink = func(k int, N int) {
		for 2*k+1 < N {
			j := 2*k + 1
			if j+1 < N && nums[j] < nums[j+1] {
				j++
			}
			if nums[k] >= nums[j] {
				break
			}
			nums[k], nums[j] = nums[j], nums[k]
			k = j
		}
	}

	for k := N/2 - 1; k >= 0; k-- {
		sink(k, N)
	}

	for N > 1 {
		nums[0], nums[N-1] = nums[N-1], nums[0]
		N--
		sink(0, N)
	}

	return nums
}
