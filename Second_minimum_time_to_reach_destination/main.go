package main

import "fmt"

func main() {
	secondMinimum(5, [][]int{{1, 2}, {1, 3}, {1, 4}, {3, 4}, {4, 5}}, 3, 5)
	secondMinimum(2, [][]int{{1, 2}}, 3, 2)
	secondMinimum(2, [][]int{{1, 2}}, 2, 1)

}

func secondMinimum(n int, edges [][]int, time int, change int) int {
	ref := make([][][]int, n)
	for r := range ref {
		ref[r] = make([][]int, n)
		for c := range ref[r] {
			ref[r][c] = make([]int, 2)
			for m := range ref[r][c] {
				ref[r][c][m] = 9999
			}
		}
	}

	for _, e := range edges {
		ref[e[0]-1][e[1]-1][0] = time
		ref[e[1]-1][e[0]-1][0] = time
	}

	for k := range n {
		for i := range n {
			for j := range n {
				addMin(&ref[i][j], ref[i][k][0]+ref[k][j][0])
				addMin(&ref[i][j], ref[i][k][1]+ref[k][j][1])
				addMin(&ref[i][j], ref[i][k][1]+ref[k][j][0])
				addMin(&ref[i][j], ref[i][k][0]+ref[k][j][1])
			}
		}
	}

	fmt.Println(ref)
	status := true
	steps := 0
	newChange := change
	answer := ref[0][n-1][1]
	for {
		if status {
			for range time {
				steps++
				if answer == steps {
					fmt.Println(answer)
					return answer
				}
				if steps == newChange {
					status = !status
					newChange += change
				}
			}
		} else {
			for steps < newChange {
				steps++
				answer++
			}
			status = !status
			newChange += change
		}
	}
}

func addMin(minList *[]int, num int) {
	if num < (*minList)[0] {
		(*minList)[0], (*minList)[1] = (*minList)[1], (*minList)[0]
		(*minList)[0] = num
	} else if num < (*minList)[1] {
		(*minList)[1] = num
	}
}
