package main

func main() {
	findTheCity(5, [][]int{{0, 1, 2}, {0, 4, 8}, {1, 2, 3}, {1, 4, 2}, {2, 3, 1}, {3, 4, 1}}, 2)
}

func findTheCity(n int, edges [][]int, distanceThreshold int) int {
	ref := make([][]int, n)
	for r := range ref {
		ref[r] = make([]int, n)
		for c := range ref[r] {
			ref[r][c] = 99999
		}
	}

	for _, e := range edges {
		if e[2] <= distanceThreshold {
			ref[e[0]][e[1]] = e[2]
		}
	}

	for k := range n {
		for i := range n {
			for j := i + 1; j < n; j++ {
				if min(ref[i][j], ref[i][k]+ref[k][j]) <= distanceThreshold {
					ref[i][j] = min(ref[i][j], ref[i][k]+ref[k][j])
					ref[j][i] = ref[i][j]
				}
			}
		}
	}
	max := 0
	min_count := n - 1
	for i := range n {
		count := 0
		for j := range ref[i] {
			if ref[i][j] != 99999 {
				count++
			}
		}
		if count <= min_count {
			max = i
			min_count = count
		}
	}

	return max
}
