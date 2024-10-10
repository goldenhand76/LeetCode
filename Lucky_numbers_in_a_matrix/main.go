package main

func main() {
	luckyNumbers([][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}})
	luckyNumbers([][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}})
	luckyNumbers([][]int{{7, 8}, {1, 2}})
}

func luckyNumbers(matrix [][]int) []int {
	if len(matrix) == 0 {
		return nil
	}

	rowMins := make([]int, len(matrix))
	colMaxs := make([]int, len(matrix[0]))

	for i := range matrix {
		rowMins[i] = matrix[i][0]
		for j := range matrix[i] {
			if matrix[i][j] < rowMins[i] {
				rowMins[i] = matrix[i][j]
			}
		}
	}

	for j := range matrix[0] {
		colMaxs[j] = matrix[0][j]
		for i := range matrix {
			if matrix[i][j] > colMaxs[j] {
				colMaxs[j] = matrix[i][j]
			}
		}
	}

	result := []int{}
	for i := range rowMins {
		for j := range colMaxs {
			if rowMins[i] == colMaxs[j] {
				result = append(result, rowMins[i])
			}
		}
	}

	return result
}
