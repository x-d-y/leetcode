package main

import "fmt"

func main(){
	fmt.Println(validMountainArray([]int{9,8,7,6,5,4,3,2,1}))
}
func validMountainArray(A []int) bool {
	if len(A)<3 {
		return false
	}
	start := 0
	end := len(A)-1
	startTop := false
	endTop := false
	
	if A[1]<=A[0] || A[end-1] <= A[end]{
		return false
	}
	
	for range A{
		if A[start] < A[start+1] && !startTop{
			start ++
		}else{
			startTop = true
		}
		if A[end] < A[end -1] && !endTop{
			end --
		}else{
			endTop = true
		}
		if start == end{
			return true
		}
		if startTop && endTop{
			return false
		}
	}
	return false
}