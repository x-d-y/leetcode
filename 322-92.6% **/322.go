package main

import (
	"fmt"
)

func main() {
	fmt.Print(coinChange([]int{2}, 3))
}

func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)
	for i, _ := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
