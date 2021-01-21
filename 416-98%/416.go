package main

import "fmt"

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}

func canPartition(nums []int) bool {
	n := len(nums)
	c := 0
	for i := 0; i < n; i++ {
		c += nums[i]
	}
	if c%2 != 0 {
		return false
	}
	dp := make([]bool, c+1)
	dp[0] = true
	target := c / 2
	for i := 0; i < n; i++ {
		for j := target; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
			if dp[target] == true {
				return true
			}
		}
	}
	return dp[target]
}
