package main

import "fmt"

func main() {
	r := kthDistinct([]string{"d", "b", "c", "b", "c", "b", "c", "a", "b", "c"}, 2)
	fmt.Println(r)
	r = kthDistinct([]string{"aaa", "aa", "a"}, 1)
	fmt.Println(r)
	r = kthDistinct([]string{"a", "b", "a"}, 3)
	fmt.Println(r)
	r = kthDistinct([]string{"a", "b", "a", "c"}, 3)
	fmt.Println(r)
}

func kthDistinct(arr []string, k int) string {
	m := make(map[string]int)

	// Count occurrences of each string
	for _, s := range arr {
		m[s]++
	}

	// Collect the distinct elements
	distinctElements := []string{}
	for _, s := range arr {
		if m[s] == 1 {
			distinctElements = append(distinctElements, s)
		}
	}

	// Check if there are at least k distinct elements
	if len(distinctElements) < k {
		return ""
	}

	// Return the k-th distinct element
	return distinctElements[k-1]
}
