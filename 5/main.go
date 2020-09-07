package main

func longestPalindrome(s string) string {
	if s == "" || len(s)==0{
		return ""
	}
	L := 0
	H := 0
	low := 0
	high := 0
	long := 0
	for i:=0 ;i<len(s);i++{
		low,high =lh(&i,&s)
		getMid(&low,&high,&s)
		if high-low >long{
			H = high-1
			L = low+1
			long = high-low
		}
	}
	return s[L:H+1]
}

func lh(i *int, s *string)(int,int){
	low := *i
	high := *i
	for {
		if *i >= 0 && *i+1 < len(*s) && (*s)[*i] == (*s)[*i+1] {
			high++
			*i++
		}else{
			break
		}

	}
	*i = high
	return low,high
}

func getMid(low ,high *int,s *string) {
	for *low>=0 && *high<len(*s){
		if (*s)[*low] == (*s)[*high]{
			*low--
			*high++
		}else{
			return
		}
	}
}


//func main(){
//	s := "babad"
//	fmt.Println(longestPalindrome(s))
//}
