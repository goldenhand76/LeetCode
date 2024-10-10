package main

import "fmt"

func main() {
	ar := []int{180, 165, 170}
	names := []string{"a", "b", "c"}
	sort(&ar, &names, 0, len(ar)-1)
	fmt.Println(ar)
	fmt.Println(names)
}

func partition(a *[]int, n *[]string, lo int, hi int) int {
	i := lo + 1
	j := hi

	for {
		for (*a)[i] < (*a)[lo] {
			if i == hi {
				break
			}
			i++
		}
		for (*a)[lo] < (*a)[j] {
			if j == lo {
				break
			}
			j--
		}
		if i >= j {
			break
		}
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
		(*n)[i], (*n)[j] = (*n)[j], (*n)[i]
	}
	(*a)[lo], (*a)[j] = (*a)[j], (*a)[lo]
	(*n)[lo], (*n)[j] = (*n)[j], (*n)[lo]
	return j
}

func sort(a *[]int, n *[]string, lo int, hi int) {
	if hi <= lo {
		return
	}
	j := partition(a, n, lo, hi)
	sort(a, n, lo, j-1)
	sort(a, n, j+1, hi)
}
