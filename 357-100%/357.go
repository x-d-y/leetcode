package main

import "fmt"

func main() {
	fmt.Println(countNumbersWithUniqueDigits(2))
}

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	cLeft := 9
	total := 10
	cRight := 9
	for i := 2; i <= n; i++ {
		cRight = cRight * cLeft
		total += cRight
		cLeft--
	}
	return total
}
