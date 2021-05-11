package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestValidParentheses(")()())"))
}

func longestValidParentheses(s string) int {
	stack, res := []int{}, 0
	stack = append(stack, -1)
	for i, v := range s {
		if string(v) == "(" {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	} else {
		return j
	}
}
