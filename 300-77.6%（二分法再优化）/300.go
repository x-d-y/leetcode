package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(lengthOfLIS([]int{11, 12, 13, 14, 15, 6, 7, 8, 101, 18}))
}

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, 2)
	dp[0] = 0 - int(math.Pow10(4)) - 1
	dp[1] = nums[0]
	for _, v := range nums {
		if v > dp[len(dp)-1] {
			dp = append(dp, v)
			continue
		}
		maxJ := len(dp)
		minJ := 0
		for j := len(dp) / 2; ; {
			if v > dp[j-1] {
				if v <= dp[j] {
					dp[j] = v
					break
				} else {
					minJ = j
					j = (j + maxJ) / 2
					continue
				}
			} else {
				maxJ = j
				j = (j + minJ) / 2
			}
		}
	}
	return len(dp) - 1
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
