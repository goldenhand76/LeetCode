package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	numMagicSquaresInside([][]int{
		{4, 3, 8, 4},
		{9, 5, 1, 9},
		{2, 7, 6, 2},
	})
}

// Function to count the number of 3x3 magic squares in the grid
func numMagicSquaresInside(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[i])-2; j++ {
			if isMagicSquare(grid, i, j) {
				count++
			}
		}
	}
	fmt.Println(count)
	return count
}

// Helper function to check if the 3x3 subgrid starting at grid[i][j] is a magic square
func isMagicSquare(grid [][]int, i, j int) bool {
	// Flatten the 3x3 grid into a single array
	nums := []int{
		grid[i][j], grid[i][j+1], grid[i][j+2],
		grid[i+1][j], grid[i+1][j+1], grid[i+1][j+2],
		grid[i+2][j], grid[i+2][j+1], grid[i+2][j+2],
	}

	// Check if the numbers are 1 to 9 without repetition
	used := make([]bool, 10)
	for _, num := range nums {
		if num < 1 || num > 9 || used[num] {
			return false
		}
		used[num] = true
	}

	// Check rows, columns, and diagonals
	return nums[0]+nums[1]+nums[2] == 15 && // Row 1
		nums[3]+nums[4]+nums[5] == 15 && // Row 2
		nums[6]+nums[7]+nums[8] == 15 && // Row 3
		nums[0]+nums[3]+nums[6] == 15 && // Column 1
		nums[1]+nums[4]+nums[7] == 15 && // Column 2
		nums[2]+nums[5]+nums[8] == 15 && // Column 3
		nums[0]+nums[4]+nums[8] == 15 && // Diagonal 1
		nums[2]+nums[4]+nums[6] == 15 // Diagonal 2
}

func init() {
	debug.SetGCPercent(-1)
}
