package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(11111111111111111111111111111101))
}
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		num = num & (num - 1)
		count++
	}
	return count
}
