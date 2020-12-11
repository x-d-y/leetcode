package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	Max, Min, res := 0, 0, 0

	for _, v := range nums {
		if v < 0 {
			Max, Min = Min, Max
		}
		Max = max(Max*v, v)
		Min = min(Min*v, v)
		res = max(Max, res)
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
