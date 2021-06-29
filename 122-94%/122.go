package main

import "fmt"

func main() {
	a := []int{0}
	fmt.Println(maxProfit(a))
}

var max int

func maxProfit(prices []int) int {
	max = 0
	start := 0
	for i := 0; i < len(prices)-1; i++ {
		if prices[i+1] > prices[i] {
			continue
		} else {
			max += prices[i] - prices[start]
			start = i + 1
		}
	}
	if prices[len(prices)-1] > prices[start] {
		max += prices[len(prices)-1] - prices[start]
	}
	return max
}
