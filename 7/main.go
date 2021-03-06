package main

//func main(){
//	r := reverse(120-95%)
//	fmt.Println(r)
//}

func reverse(x int) int {
	var n int64
	for x != 0 {
		n = n*10 + int64(x%10)
		x = x / 10
	}
	if n > 1<<31-1 || n < -0-1<<31 {
		return 0
	}
	return int(n)
}
