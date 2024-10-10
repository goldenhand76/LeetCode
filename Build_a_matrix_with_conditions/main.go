package main

import "fmt"

func main() {
	ans := buildMatrix(3, [][]int{{1, 2}, {3, 2}}, [][]int{{2, 1}, {3, 2}})
	fmt.Println(ans)

}

func buildMatrix(k int, rowConditions [][]int, colConditions [][]int) [][]int {
	row := make([]int, k)
	col := make([]int, k)

	for _, r := range rowConditions {
		if row[0] == 0 && row[1] == 0 {
			row[0] = r[0]
			row[1] = r[1]
		} else {
			row = find(row, r[0], r[1])[:k]
		}
	}

	for _, c := range colConditions {
		if col[0] == 0 && col[1] == 0 {
			col[0] = c[0]
			col[1] = c[1]
		} else {
			col = find(col, c[0], c[1])[:k]
		}
	}

	ans := make([][]int, k)
	for i := range ans {
		ans[i] = make([]int, k)
	}

	for i := range row {
		for j := range col {
			if row[i] == col[j] {
				ans[i][j] = row[i]
			}
		}
	}
	return ans
}

func find(list []int, n int, lb int) []int {
	for i, v := range list {
		if v == 0 {
			list[i] = n
			return list
		}
		if v == lb {
			list = append(list[:i+1], list[i:]...)
			list[i] = n
			return list
		}
	}
	return []int{}
}
