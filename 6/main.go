package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "LEETCODEISHIRING"
	fmt.Println(convert(s, 3))
}

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	dir := false
	var re []string
	cuRow := 0

	for i := 0; i < numRows; i++ {
		re = append(re, "")
	}

	for _, v := range s {
		re[cuRow] = re[cuRow] + string(v)
		if cuRow == 0 || cuRow == numRows-1 {
			dir = !dir
		}
		if dir == true {
			cuRow++
			continue
		}
		cuRow--

	}
	return strings.Join(re, "")
}

//func convert(s string, numRows int) string {
//	if numRows == 1 {
//		return s
//	}
//	re := ""
//	//fmt.Println(len(s) / (2 * (numRows - 1)))
//	for cn := 0; cn < numRows; cn++ {
//		for i := 0; i <= len(s)/(2*(numRows-1)); i++ {
//			if cn == 0 || cn == numRows-1 {
//				if 2*(numRows-1)*i+cn < len(s) {
//					//fmt.Println(string(s[2*(numRows-1)*i+cn]))
//					re += string(s[2*(numRows-1)*i+cn])
//				}
//			} else {
//				if 2*(numRows-1)*i+cn < len(s) {
//					//fmt.Println(string(s[2*(numRows-1)*i+cn]))
//					re += string(s[2*(numRows-1)*i+cn])
//				}
//				if 2*(numRows-1)*i+cn+2*(numRows-cn-1) < len(s) {
//					//fmt.Println(string(s[2*(numRows-1)*i+cn+2*(numRows-cn-1)]))
//					re += string(s[2*(numRows-1)*i+cn+2*(numRows-cn-1)])
//				}
//			}
//		}
//	}
//	return re
//}
