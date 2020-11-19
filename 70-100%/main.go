package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
}

func climbStairs(n int) int {

	if n == 1 {
		return 1
	}
	dp := []int{1, 1}

	for i := 2; i < n+1; i++ {
		dp = append(dp, dp[i-1]+dp[i-2])
	}
	return dp[n]
}
