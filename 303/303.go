package main

import "fmt"

func main() {
	num := []int{-2, 0, 3, -5, 2, -1}
	st := segmentTree(num, 0, 0, len(num)-1)
	fmt.Println(result(st, 0, 5))
}

func result(s *st, start, end int) int {
	sum := 0
	if start == s.start && end == s.end {
		return s.sum
	}
	if s.left.end >= end {
		sum += result(s.left, start, end)
	} else if s.right.start <= start {
		sum += result(s.right, start, end)
	} else {
		sum += result(s.left, start, s.left.end)
		sum += result(s.right, s.right.start, end)
	}
	return sum
}

type st struct {
	index, start, end int
	left              *st
	right             *st
	sum               int
}

func segmentTree(num []int, index, left, right int) *st {
	st := new(st)
	st.index = index
	st.start = left
	st.end = right

	if st.start == st.end {
		st.sum = num[st.start]
		return st
	}

	iLeft, iRight := index*2+1, index*2+2
	mid := (left + right) / 2
	st.left = segmentTree(num, iLeft, left, mid)
	st.right = segmentTree(num, iRight, mid+1, right)

	st.sum = st.left.sum + st.right.sum
	return st
}
