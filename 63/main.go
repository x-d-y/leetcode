package main

import "fmt"

func main() {
	fmt.Println(uniquePathsWithObstacles([][]int{[]int{0, 0, 0}, []int{0, 1, 0}, []int{0, 0, 0}}))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {

	var cRec [][]int

	lr := len(obstacleGrid)
	lc := len(obstacleGrid[0])

	for i := 0; i < lr; i++ {
		c := make([]int, lc)
		cRec = append(cRec, c)
	}

	var row, col int
	cRec[0][0] = 1
	for rI, rV := range obstacleGrid {
		for cI, _ := range rV {
			if obstacleGrid[rI][cI] == 1 {
				cRec[rI][cI] = 0
				continue
			}
			if rI-1 < 0 {
				row = 0
			} else {
				row = cRec[rI-1][cI]
			}
			if cI-1 < 0 {
				col = 0
			} else {
				col = cRec[rI][cI-1]
			}
			cRec[rI][cI] = row + col
			if rI == cI && cI == 0 {
				cRec[rI][cI] = 1
			}
		}
	}
	return cRec[lr-1][lc-1]
}

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	var rec [][]int

	for i := 0; i < m; i++ {
		c := make([]int, n)
		rec = append(rec, c)
	}
	rec[0][0] = 1
	var row, col int
	for rI, rV := range rec {
		for cI, _ := range rV {
			if rI-1 < 0 {
				row = 0
			} else {
				row = rec[rI-1][cI]
			}
			if cI-1 < 0 {
				col = 0
			} else {
				col = rec[rI][cI-1]
			}
			rec[rI][cI] = row + col
			if rI == cI && cI == 0 {
				rec[rI][cI] = 1
			}
		}
	}

	return rec[m-1][n-1]

}
