package main

import "fmt"

func main() {
	result := threeConsecutiveOdds([]int{1, 2, 1, 1})
	fmt.Println(result)
}

func threeConsecutiveOdds(arr []int) bool {
	count := 0
	for i, v := range arr {
		switch len(arr) - i + 1 {
		case 1:
			if count < 2 {
				return false
			}
		case 2:
			if count == 0 {
				return false
			}
		}

		if v%2 != 0 {
			count += 1
		} else {
			count = 0
		}
		if count == 3 {
			return true
		}
	}
	return false
}
