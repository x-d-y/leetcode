package main

import (
	"fmt"
	"strconv"
)


func main(){
	root := CreateNodeTree([]string{"null"})
	fmt.Println(inorderTraversal(root))

}

func inorderTraversal(root *TreeNode) []int {

	var arr []int
	r(root,&arr)
	return arr
}

func r(root *TreeNode,arr *[]int){
	if root == nil{
		return
	}
	if root.Left != nil{
		r(root.Left,arr)
	}
	*arr = append(*arr,root.Val)
	if root.Right != nil{
		r(root.Right,arr)
	}
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

