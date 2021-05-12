package main

import "fmt"

func main() {
	fmt.Println(trap([]int{4, 2, 0, 3, 2, 5}))
}

const (
	rightDir = 1
	leftDir  = -1
)

var v int

func trap(height []int) int {
	v = 0
	left, right := 0, len(height)-1

	if height[left] <= height[right] {
		scan(left, right, rightDir, height)
	} else {
		scan(right, left, leftDir, height)
	}

	return v
}

func scan(start, end, dir int, height []int) {
	if start == end {
		return
	}
	for i := start; i != end; {
		i += dir
		if height[i] >= height[start] {
			if height[i] <= height[end] {
				scan(i, end, rightDir*dir, height)
			} else {
				scan(end, i, leftDir*dir, height)
			}
			break
		}
		v += height[start] - height[i]
	}
}
