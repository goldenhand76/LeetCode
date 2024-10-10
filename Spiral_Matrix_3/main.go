package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	var r [][]int
	r = spiralMatrixIII(1, 4, 0, 0)
	fmt.Println(r)
}

func spiralMatrixIII(rows int, cols int, rStart int, cStart int) [][]int {
	answer := [][]int{{rStart, cStart}}
	pointer := []int{rStart, cStart}
	unit := 1
	totalCells := rows * cols

	goTo := func(direction string, unit int) {
		switch direction {
		case "Right":
			for i := 1; i <= unit; i++ {
				pointer[1]++
				if pointer[1] < cols && pointer[1] >= 0 && pointer[0] >= 0 && pointer[0] < rows {
					answer = append(answer, []int{pointer[0], pointer[1]})
				}
			}
		case "Down":
			for i := 1; i <= unit; i++ {
				pointer[0]++
				if pointer[0] < rows && pointer[0] >= 0 && pointer[1] >= 0 && pointer[1] < cols {
					answer = append(answer, []int{pointer[0], pointer[1]})
				}
			}
		case "Left":
			for i := 1; i <= unit; i++ {
				pointer[1]--
				if pointer[1] >= 0 && pointer[1] < cols && pointer[0] >= 0 && pointer[0] < rows {
					answer = append(answer, []int{pointer[0], pointer[1]})
				}
			}
		case "Up":
			for i := 1; i <= unit; i++ {
				pointer[0]--
				if pointer[0] >= 0 && pointer[0] < rows && pointer[1] >= 0 && pointer[1] < cols {
					answer = append(answer, []int{pointer[0], pointer[1]})
				}
			}
		}
	}

	for len(answer) < totalCells {
		goTo("Right", unit)
		goTo("Down", unit)
		unit++
		goTo("Left", unit)
		goTo("Up", unit)
		unit++
	}

	return answer
}

func init() {
	debug.SetGCPercent(-1)
}
