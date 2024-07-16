package main

import (
	"fmt"
)

// TODO: The code is not yet completed
func main() {
	maxNumEdgesToRemove(4, [][]int{{3, 1, 2}, {3, 2, 3}, {1, 1, 3}, {1, 2, 4}, {1, 1, 2}, {2, 3, 4}})
}

func maxNumEdgesToRemove(n int, edges [][]int) int {
	alice := make(map[int][]int)
	bob := make(map[int][]int)
	both := make(map[int][]int)
	for _, v := range edges {
		if v[0] == 1 {
			alice[v[1]] = append(alice[v[1]], v[2])
		} else if v[0] == 2 {
			bob[v[1]] = append(bob[v[1]], v[2])
		} else {
			both[v[1]] = append(both[v[1]], v[2])
		}
	}
	for k, v1 := range both {
		for v2 := range v1 {
			if v3, ok := both[v2]; ok {
				for v4 := range v3 {
					both[k] = append(both[k], v4)
				}
			}
		}
	}
	fmt.Println("n:", n)
	fmt.Println("alice:", alice)
	fmt.Println("bob:", bob)
	fmt.Println("both:", both)
	return -1
}
