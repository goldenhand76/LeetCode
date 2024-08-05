package main

import "fmt"

func main() {
	fmt.Println(canBeEqual([]int{1, 2, 3, 4}, []int{2, 4, 1, 3}))
	fmt.Println(canBeEqual([]int{7}, []int{7}))
	fmt.Println(canBeEqual([]int{3, 7, 9}, []int{3, 7, 11}))
}

func canBeEqual(target []int, arr []int) bool {
	keys := make(map[int]int)
	for i := range target {
		if _, ok := keys[target[i]]; !ok {
			keys[target[i]] = 1
		} else {
			keys[target[i]] += 1
		}
	}
	for j := range arr {
		if _, ok := keys[arr[j]]; ok {
			keys[arr[j]] -= 1
			if keys[arr[j]] < 0 {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
