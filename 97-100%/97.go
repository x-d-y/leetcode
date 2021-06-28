package main

import "fmt"

func main() {
	s1 := "abababababababababababababababababababababababababababababababababababababababababababababababababbb"
	s2 := "babababababababababababababababababababababababababababababababababababababababababababababababaaaba"
	s3 := "abababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababbb"

	fmt.Println(isInterleave(s1, s2, s3))

}

// 深度优先

var gs1, gs2, gs3 string

func isInterleave(s1, s2, s3 string) bool {
	if len(s3) != (len(s1) + len(s2)) {
		return false
	}
	gs1, gs2, gs3 = s1, s2, s3

	u := make(map[int]bool)
	return re(0, 0, 0, u)

}

func re(i, x, y int, u map[int]bool) bool {
	if i == len(gs3) {
		return true
	}

	if _, ok := u[(x)*len(gs3)+y]; ok {
		return false
	}

	u[(x)*len(gs3)+y] = true


	if x < len(gs1) && gs3[i] == gs1[x] {

		if re(i+1, x+1, y, u) {
			return true
		}
	}

	if y < len(gs2) && gs3[i] == gs2[y] {
		if re(i+1, x, y+1, u) {
			return true
		}
	}

	return false
}

// 广度优先

// func isInterleave(s1 string, s2 string, s3 string) bool {
// 	if len(s3) != (len(s1) + len(s2)) {
// 		return false
// 	}
// 	var x int
// 	var p [][]int
// 	p = append(p, make([]int, 0), make([]int, 0))
// 	p[0] = append(p[0], 0)
// 	for i := 0; i < len(s3); i++ {
// 		p[(i+1)%2] = make([]int, 0)
// 		for j := 0; j < len(p[i%2]); j++ {
// 			x = p[i%2][j]
// 			if x < len(s1) && s3[i] == s1[x] {
// 				p[(i+1)%2] = append(p[(i+1)%2], x+1)
// 			}
// 			if i-x < len(s2) && s3[i] == s2[i-x] {
// 				p[(i+1)%2] = append(p[(i+1)%2], x)
// 			}
// 		}
// 	// 	if len(p[(i+1)%2]) == 0 {
// 	// 		return false
// 	// 	}
// 	// }
// 	return true
// }
