package utils

import (
	"fmt"
	"sudoku/common"
)

const N = 9

func BuildSudokuRequest() common.SudokuRequest {
	grid := [common.N][common.N]int{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	return common.SudokuRequest{
		Grid: grid,
	}
}
func isSafe(grid *[N][N]int, row, col, num int) bool {
	for x := 0; x < N; x++ {
		if grid[row][x] == num {
			return false
		}
	}

	for x := 0; x < N; x++ {
		if grid[x][col] == num {
			return false
		}
	}

	startRow := row - row%3
	startCol := col - col%3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[i+startRow][j+startCol] == num {
				return false
			}
		}
	}

	return true
}

func SolveSudoku(grid *[N][N]int) bool {
	var row, col int
	empty := false
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if grid[i][j] == 0 {
				row, col = i, j
				empty = true
				break
			}
		}
		if empty {
			break
		}
	}

	if !empty {
		return true
	}

	for num := 1; num <= 9; num++ {
		if isSafe(grid, row, col, num) {
			grid[row][col] = num
			if SolveSudoku(grid) {
				return true
			}
			grid[row][col] = 0
		}
	}
	return false
}

func PrintGrid(grid [N][N]int) {
	for r := 0; r < N; r++ {
		for d := 0; d < N; d++ {
			fmt.Printf("%d ", grid[r][d])
		}
		fmt.Printf("\n")
	}

}
