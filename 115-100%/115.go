package main

import "fmt"

func main() {

	s := "rabbbit"
	t := "rabbit"
	fmt.Println(numDistinct(s, t))

}

func numDistinct(s string, t string) int {
	if len(s) < len(t) {
		return 0
	}
	dp := make([]int, len(t)+1)
	dp[len(t)] = 1
	for i := len(s) - 1; i >= 0; i-- {
		right := dp[len(t)]
		for j := len(t) - 1; j >= 0; j-- {
			if i > j+len(s)-len(t) {
				break
			}
			left := dp[j]
			if s[i] == t[j] {
				dp[j] = right + left
			} else {
				dp[j] = left
			}
			right = left
		}

	}
	return dp[0]
}
