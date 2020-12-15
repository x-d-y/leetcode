package main

import "fmt"

func main() {
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	last2, last := 0, nums[0]
	for i := 1; i < len(nums); i++ {
		last, last2 = max(last2+nums[i], last), last
	}
	return max(last, last2)
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
