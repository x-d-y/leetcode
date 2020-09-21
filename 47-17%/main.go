package main

import (
	"fmt"
	"sort"
)

//func main(){
//	permuteUnique([]int{2,2,1,1})
//}

var mat [][]int
var records map[int]bool
func permuteUnique(nums []int) [][]int {
	mat = [][]int{}
	vis := make([]bool,len(nums))
	records = make(map[int]bool)
	sort.Ints(nums)
	r(nums,vis,[]int{})
	fmt.Println(mat)
	return mat
}

func r(nums []int,vis []bool,cu []int){
	for i,v := range nums{
		if vis[i] || i>0 && v== nums[i-1] &&!vis[i-1]{
			continue
		}
		cu = append(cu, v)
		vis[i] = true

		if len(cu) == len(nums){
			mat = append(mat, append([]int(nil), cu...))
			mat = append(mat,cu)
			vis[i] = false

			return
		}
		r(nums,vis,cu)
		cu = cu[:len(cu)-1]
		vis[i] = false
	}
}
