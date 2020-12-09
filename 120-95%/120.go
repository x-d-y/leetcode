package main

import "fmt"

func main() {
	t := [][]int{[]int{2}, []int{3, 4}, []int{5, 6, 7}, []int{1, 5, 3, 6}}
	fmt.Println(minimumTotal(t))
}

func minimumTotal(triangle [][]int) int {
	res := triangle[0][0]
	if len(triangle) == 1 {
		return res
	}

	for i := 1; i < len(triangle); i++ {
		for ii, _ := range triangle[i] {
			if ii == 0 {
				triangle[i][ii] += triangle[i-1][ii]
				res = triangle[i][ii]
				continue
			}
			if ii == i {
				triangle[i][ii] += triangle[i-1][ii-1]
				if res > triangle[i][ii] {
					res = triangle[i][ii]
				}
				continue
			}

			triangle[i][ii] += min(triangle[i-1][ii], triangle[i-1][ii-1])
			if res > triangle[i][ii] {
				res = triangle[i][ii]
			}
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}
