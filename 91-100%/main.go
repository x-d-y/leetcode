package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(numDecodings("01"))
}

func numDecodings(s string) int {
	dp := []int{1, 1}
	if s[0] == 48 {
		return 0
	}
	last := 1
	for i := 1; i < len(s); i++ {

		dp = append(dp, dp[i])

		c2, _ := strconv.Atoi(s[i-1 : i+1])
		if c2 == 0 || c2 > 26 && c2%10 == 0 {
			return 0
		}
		if c2 <= 26 && c2%10 != 0 && last != 0 {
			dp[i+1] += dp[i-1]
		}
		last, _ = strconv.Atoi(s[i : i+1])
		if last == 0 {
			dp[i+1] = dp[i-1]
		}
	}
	return dp[len(s)]
}
