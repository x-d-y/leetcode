package main

import "fmt"

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func maxProfit(prices []int) int {

	fmt.Println(prices[0:0])
	if len(prices) <= 1 {
		return 0
	}
	maxProfit := 0
	min := prices[0]
	for _, v := range prices {
		if v-min > maxProfit {
			maxProfit = v - min
		}

		if v < min {
			min = v
		}

	}
	return maxProfit
}
