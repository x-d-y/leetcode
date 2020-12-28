package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}

func maxProfit_(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	sell := make([]int, len(prices)+1)
	buy := make([]int, len(prices)+1)
	sell[0], sell[1] = 0, 0
	buy[0], buy[1] = math.MinInt64, -prices[0]
	for i := 2; i <= len(prices); i++ {
		buy[i] = max(buy[i-1], sell[i-2]-prices[i-1])
		sell[i] = max(sell[i-1], prices[i-1]+buy[i-1])
	}
	return sell[len(prices)]
}

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	sell_1, sell_2 := 0, 0
	buy_0, buy_1 := math.MinInt64, -prices[0]
	for i := 2; i <= len(prices); i++ {
		buy_1, buy_0 = max(buy_1, sell_1-prices[i-1]), buy_1
		sell_2, sell_1 = max(sell_2, prices[i-1]+buy_0), sell_2
	}
	return sell_2
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
