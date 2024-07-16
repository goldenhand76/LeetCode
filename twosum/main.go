package main

import "fmt"

func main() {
	result := twoSum([]int{1, 2, 3, 4, 5}, 6)
	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]*int)
	for i := 0; i < len(nums); i++ {
		if val, ok := m[target-nums[i]]; ok {
			return []int{i, *val}
		}
		m[nums[i]] = &i
	}
	return []int{}
}
