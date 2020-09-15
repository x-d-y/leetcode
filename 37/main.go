package main

import (
	"fmt"
)

func main(){

	b := [][]byte{
		[]byte{'5','3','.','.','7','.','.','.','.'},
		[]byte{'6','.','.','1','9','5','.','.','.'},
		[]byte{'.','9','8','.','.','.','.','6','.'},
		[]byte{'8','.','.','.','6','.','.','.','3'},
		[]byte{'4','.','.','8','.','3','.','.','1'},
		[]byte{'7','.','.','.','2','.','.','.','6'},
		[]byte{'.','6','.','.','.','.','2','8','.'},
		[]byte{'.','.','.','4','1','9','.','.','5'},
		[]byte{'.','.','.','.','8','.','.','7','9'},
	}
	solveSudoku(b)
	fmt.Println(b)
}


var dot [][2]int
func solveSudoku(board [][]byte)  {
	dot = [][2]int{}
	var row [9][9]bool
	var column [9][9]bool
	var matrix [3][3][9]bool

	for i,r := range board{
		for j,c := range r{
			if c == '.'{
				dot = append(dot,[2]int{i,j})
			}else{
				row[i][c-'1'] = true
				column[j][c-'1'] = true
				matrix[i/3][j/3][c-'1'] = true
			}
		}
	}
	//fmt.Println(board)
	if r(0,row,column,matrix,board){
		return
	}

}

func r(d int,row,column [9][9]bool, matrix [3][3][9]bool,b [][]byte)bool{
	if len(dot)==0{
		return true
	}
	for i:= 0; i<9; i++{
		rw,c := dot[d][0], dot[d][1]
		if row[rw][i] == false&&column[c][i] == false&&matrix[rw/3][c/3][i] == false {
			row[rw][i] = true
			column[c][i] = true
			matrix[rw/3][c/3][i] = true
			b[rw][c] = byte(i+'1')
			if len(dot)-1 ==d{
				return true
			}
			if r(d+1,row,column,matrix,b){
				return true
			}
			row[rw][i] = false
			column[c][i] = false
			matrix[rw/3][c/3][i] = false
		}
	}
	return false
}