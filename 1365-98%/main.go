package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	count := [101]int{}
	for _,v := range nums{
		count[v]++
	}
	var cCount int
	var cNum int
	for i,v := range count{
		cNum = v
		count[i] = cCount
		cCount += cNum
	}
	for i,v := range nums{
		nums[i] = count[v]
	}
	return nums
}

func main(){
	fmt.Println(smallerNumbersThanCurrent([]int{37,64,63,2,41,78,51,36,2,20,25,41,72,100,17,43,54,27,34,86,12,48,70,44,87,68,62,98,68,30,8,92,5,10}))
}