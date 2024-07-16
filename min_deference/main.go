package main

import "fmt"

func main() {
	fmt.Println(minDifference([]int{5, 3, 2, 4}))
	fmt.Println(minDifference([]int{1, 5, 0, 10, 14}))
	fmt.Println(minDifference([]int{3, 100, 20}))
}

func minDifference(nums []int) int {
	if len(nums) < 5 {
		return 0
	}
	for i := 1; i < len(nums); i++ {
		v := nums[i]
		index := i
		for index > 0 && nums[index-1] > v {
			nums[index] = nums[index-1]
			index = index - 1
		}
		nums[index] = v
	}
	firstMean := (nums[0] + nums[1] + nums[2]) / 3
	lastMean := (nums[len(nums)-1] + nums[len(nums)-2] + nums[len(nums)-3]) / 3
	if len(nums) == 5 {
		if firstMean > lastMean {
			if nums[len(nums)-1]-nums[len(nums)-2] < nums[len(nums)-2]-nums[len(nums)-3] {
				return nums[len(nums)-1] - nums[len(nums)-2]
			} else {
				return nums[len(nums)-2] - nums[len(nums)-3]
			}
		} else {
			if nums[1]-nums[0] < nums[2]-nums[1] {
				return nums[1] - nums[0]
			} else {
				return nums[2] - nums[1]
			}
		}
	}
	if firstMean > lastMean {
		return nums[len(nums)-1] - nums[len(nums)-len(nums)/2]
	} else {
		return nums[len(nums)/2] - nums[0]
	}
}
