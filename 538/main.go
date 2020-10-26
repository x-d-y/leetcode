package main

import (
	"strconv"
)

func main(){
	root := CreateNodeTree([]string{"5","2","13"})
	convertBST(root)
}

var ga []*TreeNode
func convertBST(root *TreeNode) *TreeNode {
	ga = []*TreeNode{}
	r(root)
	var sum int
	for i:= len(ga)-1; i>=0;i--{
		ga[i].Val += sum
		sum = ga[i].Val
	}
	return root
}

func r(root * TreeNode){
	if root == nil{
		return
	}
	r(root.Left)
	ga = append(ga,root)
	r(root.Right)

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



