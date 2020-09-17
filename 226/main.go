package main

import (
	"fmt"
	"strconv"
)

func main(){
	root := CreateNodeTree([]string{"4","2","7","1","3","6","9"})
	invertTree(root)
	fmt.Println(root)
}


func invertTree(root *TreeNode) *TreeNode {
	if root == nil{
		return root
	}
	r(root)
	return root
}

func r(root *TreeNode){


	if root.Left == nil && root.Right == nil{
		return
	}

	if root.Left != nil{
		r(root.Left)
	}
	if root.Right != nil{
		r(root.Right)
	}
	root.Left ,root.Right = root.Right,root.Left
	return
}


type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func CreateNodeTree(a []string)*TreeNode{
	var node []*TreeNode
	left := 0
	n ,_:= strconv.Atoi(a[0])
	root := &TreeNode{Val: n}
	node = append(node,root)
	dir := false
	for right:=1;right<len(a);right++{
		c := new(TreeNode)
		n ,err:= strconv.Atoi(a[right])
		if err != nil{
			c = nil
		}else {
			c.Val = n
		}
		node = append(node,c)
		if node[left] == nil{
			left++
			dir = false
		}
		if !dir{
			node[left].Left = c
		}else{
			node[left].Right = c
			left++
		}
		dir = !(dir||false)
	}
	return root
}


