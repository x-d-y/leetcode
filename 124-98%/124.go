package main

import (
	"fmt"
	"math"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	n := makeTree([]string{"12", "3", "4", "5", "6", "7", "null", "null", "3", "4"})
	fmt.Println(maxPathSum(n))
}

func makeTree(list []string) *TreeNode {

	var node []*TreeNode
	for _, v := range list {
		if v != "null" {
			if vn, e := strconv.Atoi(v); e != nil {
				fmt.Println(e.Error())
			} else {
				n := new(TreeNode)
				n.Val = vn
				node = append(node, n)
			}
		} else {
			node = append(node, nil)
		}
	}

	for i, v := range node {
		if i*2+1 < len(node) {
			v.Left = node[i*2+1]
		}
		if i*2+2 < len(node) {
			v.Right = node[i*2+2]
		}
	}
	return node[0]
}

var maxNum int

func maxPathSum(root *TreeNode) int {
	if root.Left == nil && root.Right == nil {
		return root.Val
	}
	maxNum = math.MinInt64
	getCurMax(root)
	return maxNum
}

func getCurMax(root *TreeNode) int {

	if root == nil{
		return 0
	}
	left := getCurMax(root.Left)
	right := getCurMax(root.Right)
	cMax := max(max(left+root.Val, right+root.Val), root.Val)
	maxNum = max(maxNum, max(cMax, left+root.Val+right))
	return cMax
}
func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
