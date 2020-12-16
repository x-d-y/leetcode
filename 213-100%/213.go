package main

import "fmt"

func main() {
	fmt.Println(rob([]int{6, 3}))
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	rob1 := subRob(nums[1:])
	rob2 := subRob(nums[:len(nums)-1])
	return max(rob1, rob2)
}

func subRob(nums []int) int {
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
