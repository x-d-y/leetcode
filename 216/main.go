package main

import "fmt"

func main(){
	fmt.Println(combinationSum3(10, 50))
}

var g [][]int

func combinationSum3(k int, n int) [][]int {
	if n == 0{
		return [][]int{[]int{}}
	}
	g = [][]int{}
	var a []int
	r(1,0,k,n,a)
	return g
}

func r(cur int, sum int, num int, target int,arr []int){
	if cur+sum>target{
		return
	}
	sum += cur
	arr = append(arr,cur)
	if  len(arr)==num && sum==target{
		if sum == target {
			na := make([]int, num)
			copy(na, append(arr,cur))
			g = append(g, na)
			return
		}
		return
	}

	r(cur+1,sum,num,target, arr)
	sum = sum - cur
	arr = arr[:len(arr)-1]
	r(cur+1,sum,num,target,arr)
}
