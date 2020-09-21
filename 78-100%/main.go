package main

//func main(){
//	subsets([]int{1,2,3})
//}



func subsets(nums []int) [][]int {
	var ga [][]int
	var r func(k int, nums []int,numsInx int, cu []int)

	r = func (k int, nums []int,numsInx int, cu []int){
		if len(cu) == k{
			ga = append(ga, append([]int(nil),cu...))
			return
		}
		if numsInx==len(nums){
			return
		}

		cu = append(cu,nums[numsInx])
		r(k,nums,numsInx+1,cu)
		cu = cu[:len(cu)-1]
		r(k,nums,numsInx+1,cu)
	}
	for i:=0 ;i<= len(nums);i++{
		r(i,nums,0,[]int{})
	}
	//fmt.Println(ga)
	return ga
}



