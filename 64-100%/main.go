package main

import (
	"fmt"
)

func main() {
	fmt.Println(minPathSum([][]int{[]int{0}, []int{1}}))
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	row := len(grid)
	cRec := make([][]int, row)
	col := len(grid[0])
	for i, _ := range grid {
		cRec[i] = make([]int, col)
	}
	rows, cols := 1, 1
	for i, vr := range grid {
		for j, vc := range vr {
			if i-1 < 0 {
				rows = -1
			} else {
				rows = cRec[i-1][j]
			}
			if j-1 < 0 {
				cols = -1
			} else {
				cols = cRec[i][j-1]
			}
			cRec[i][j] = min(rows, cols) + vc
		}
	}
	return cRec[row-1][col-1]
}

func min(x, y int) int {
	if x < 0 && y >= 0 {
		return y
	}
	if y < 0 && x >= 0 {
		return x
	}
	if x < 0 && y < 0 {
		return 0
	}
	if x < y {
		return x
	}
	return y
}
