package main

func main() {
	generate(1)
}

func generate(numRows int) [][]int {
	re := make([][]int, numRows)

	for i, _ := range re {
		re[i] = make([]int, 0)
		for j := 0; j <= i; j++ {
			if j == 0 || j == i {
				re[i] = append(re[i], 1)
				continue
			}
			re[i] = append(re[i], re[i-1][j-1]+re[i-1][j])
		}
	}
	return re
}
