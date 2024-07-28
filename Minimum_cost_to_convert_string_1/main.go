package main

import (
	"fmt"
)

func main() {
	// r := minimumCost("abcd", "acbe", []byte{'a', 'b', 'c', 'c', 'e', 'd'},
	// 	[]byte{'b', 'c', 'b', 'e', 'b', 'e'}, []int{2, 5, 5, 1, 2, 20})

	// r := minimumCost("aaaa", "bbbb", []byte{'a', 'c'},
	// 	[]byte{'c', 'b'}, []int{1, 2})

	r := minimumCost("abcd", "abce", []byte{'a'},
		[]byte{'e'}, []int{1000})
	fmt.Println(r)
}

func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	a := append(original, changed...)
	am := make(map[byte]int) // Alphabet map
	index := 0
	for _, v := range a {
		if _, ok := am[v]; !ok {
			am[v] = index
			index++
		}
	}
	n := len(am)

	ref := make([][]int, n) // Alphabet Table
	for i := range ref {
		ref[i] = make([]int, n)
		for j := range ref[i] {
			ref[i][j] = 9999999
		}
	}

	for i := range cost {
		ref[am[original[i]]][am[changed[i]]] = cost[i]
	}

	for k := range n {
		for i := range n {
			for j := range n {
				ref[i][j] = min(ref[i][j], ref[i][k]+ref[k][j])
			}
		}
	}

	t := []byte(target)
	result := 0
	for i, v := range []byte(source) {
		if v != t[i] {
			if _, ok := am[v]; ok {
				if ref[am[v]][am[t[i]]] != 9999999 {
					result += ref[am[v]][am[t[i]]]
				} else {
					return -1
				}
			} else {
				return -1
			}

		}
	}
	return int64(result)
}
