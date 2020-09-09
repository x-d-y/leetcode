package main

import (
	"strconv"
)

var str []string

func binaryTreePaths(root *TreeNode) []string {
	s := ""
	str = []string{}
	re(root, s)
	//fmt.Println("")
	return str

}

func re(root *TreeNode, s string) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			str = append(str, s+strconv.Itoa(root.Val))
		} else {
			s = s + strconv.Itoa(root.Val) + "->"
			re(root.Left, s)
			re(root.Right, s)
		}
	}
	return
}
