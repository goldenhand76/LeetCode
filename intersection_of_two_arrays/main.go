package main

import "fmt"

func main() {
	fmt.Println(intersect([]int{1, 2, 2, 1}, []int{2, 2}))
}

func intersect(nums1 []int, nums2 []int) []int {
	output := []int{}
	for _, v1 := range nums1 {
		for i2, v2 := range nums2 {
			if v1 == v2 {
				output = append(output, v2)
				nums2 = append(nums2[:i2], nums2[i2+1:]...)
				break
			}
		}
	}
	return output
}
