package main

import "fmt"

func main() {
	fmt.Println(integerBreak(6))
}

func integerBreak(n int) int {
	if n <= 2 {
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 2; i <= n; i++ {
		for j := 1; j < i; j++ {
			dp[i] = max(dp[i], max(dp[i-j], i-j)*j)
		}
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
