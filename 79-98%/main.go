package main

import "fmt"

func main(){
	//board := [][]byte{
	//	[]byte{'A','B','C','E'},
	//	[]byte{'S','F','C','S'},
	//	[]byte{'A','D','E','E'}}
	board := [][]byte{
		[]byte{'a','b'},
	}
	word := "ba"
	fmt.Println(exist( board,word))
}

var vis [][]bool

var dir = [][]int{[]int{-1,0},[]int{1,0},[]int{0,-1},[]int{0,1}}
func exist(board [][]byte, word string) bool {
	if len(board)==0{
		return false
	}

	vis = [][]bool{}
	for i := range board{
		s := make([]bool,len(board[i]))
		vis = append(vis,s)
	}

	for i:=range board{
		for j:= range board[i]{
			if check(board,i,j,word,0){
				return true
			}
		}
	}
	return false
}

func check(board[][]byte,i,j int,word string,k int) bool{
	vis[i][j] = true
	defer func(){
		vis[i][j]= false
	}()
	if board[i][j] != word[k]{
		return false
	}
	if k==len(word)-1{
		return true
	}
	for _,v := range dir{
		//fmt.Println(i+v[0]>=0,i+v[0]<len(board),j+v[1]>=0,j+v[1]<len(board[i]))
		if (i+v[0]>=0 && i+v[0]<len(board) && j+v[1]>=0 && j+v[1]<len(board[i]))&&vis[i+v[0]][j+v[1]]==false  {
			if check(board,i+v[0],j+v[1],word,k+1){
				return true
			}
		}
	}

	return false
}




