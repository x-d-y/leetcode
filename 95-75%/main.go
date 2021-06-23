package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	t := generateTrees(3)
	fmt.Println(t[0].Val)
}

func generateTrees(n int) []*TreeNode {
	if n <= 0 {
		return make([]*TreeNode, 0)
	}
	return g(1,n)
}

func g(start, end int) []*TreeNode {
	t := make([]*TreeNode, 0)
	if start > end {
		t = append(t, nil)
		return t
	}
	for i := start; i <= end; i++ {
		l := g(start, i-1)
		r := g(i+1, end)
		for _, lt := range l {
			for _, rt := range r {
				node := &TreeNode{Val: i, Left: lt, Right: rt}
				t = append(t, node)
			}
		}
	}
	return t
}
