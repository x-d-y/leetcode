package main

//func main(){
//	 combine(4, 2)
//}

var array [][]int
func combine(n int, k int) [][]int {
	array = [][]int{}
	sn := 1
	var a []int
	r(sn,n,&a,k)
	return array

}

func r(c int,n int ,a *[]int,k int){
	if len(*a) + n-(c)+1 <k {
		return
	}
	if len(*a) == k{
		comb := make([]int,k)
		copy(comb,*a)
		array = append(array,comb)
		return
	}
	*a = append(*a,c)
	c = c+1
	r(c,n,a,k)
	*a = (*a)[:len(*a)-1]
	r(c,n,a,k)
}