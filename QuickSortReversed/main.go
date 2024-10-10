package main

import "fmt"

var exchange int = 0

func main() {
	a := []int{8, 1, 3, 5, 6}
	ReversedQuickSort(&exchange, &a, 0, len(a)-1)
	fmt.Println(a, "\t exchanges :", exchange)
}

func partition(exchange *int, a *[]int, lo int, hi int) int {
	mid := lo + (hi-lo)/2
	median := medianOfThree(a, lo, mid, hi)

	(*a)[lo], (*a)[median] = (*a)[median], (*a)[lo]

	i := lo + 1
	j := hi

	for {
		for (*a)[i] > (*a)[lo] {
			if i >= hi {
				break
			}
			i++
		}
		for (*a)[j] < (*a)[lo] {
			if j <= lo {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
		*exchange++
	}
	(*a)[lo], (*a)[j] = (*a)[j], (*a)[lo]
	*exchange++
	return j
}

func ReversedQuickSort(exchange *int, a *[]int, lo int, hi int) {
	if lo >= hi {
		return
	}
	j := partition(exchange, a, lo, hi)
	ReversedQuickSort(exchange, a, lo, j-1)
	ReversedQuickSort(exchange, a, j+1, hi)
}

func medianOfThree(a *[]int, lo int, mid int, hi int) int {
	if (*a)[lo] > (*a)[mid] {
		(*a)[lo], (*a)[mid] = (*a)[mid], (*a)[lo]
	}
	if (*a)[lo] > (*a)[hi] {
		(*a)[lo], (*a)[hi] = (*a)[hi], (*a)[lo]
	}
	if (*a)[mid] > (*a)[hi] {
		(*a)[mid], (*a)[hi] = (*a)[hi], (*a)[mid]
	}
	return mid
}
