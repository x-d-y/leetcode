package main

import "fmt"

func main() {
	fmt.Println(uniquePaths(0, 1))

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
