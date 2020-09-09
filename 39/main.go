package main

import "fmt"

func main(){
	a := combinationSum([]int{2,3,5}, 8)
	fmt.Println(a)
}

var a [][]int

func  combinationSum(candidates []int, target int) [][]int {
	a = [][]int{}
	if target == 0{
		return [][]int{[]int{}}
	}
	r(0,candidates,target,0,[]int{})
	return a
}

func r(c int,candidates []int,target int ,sum int,cur []int){
	if sum >target||c>len(candidates)-1{
		return
	}

	if sum == target{
		co := make([]int,len(cur))
		copy(co,cur)
		a = append(a,co)
		return
	}

	sum = sum+candidates[c]
	cur = append(cur,candidates[c])
	r(c,candidates,target,sum,cur)
	sum = sum - candidates[c]
	cur = cur[:len(cur)-1]

	r(c+1,candidates,target,sum,cur)

}
