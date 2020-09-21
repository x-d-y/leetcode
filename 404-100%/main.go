package main

import (
	"strconv"
)

//func main(){
//	root := CreateNodeTree([]string{"1","2","3","4","5"})
//	fmt.Println(sumOfLeftLeaves(root))
//
//}

func sumOfLeftLeaves(root *TreeNode) int {
	return r(root,0)
}


func r(root *TreeNode,sum int) int{
	if root == nil{
		return sum
	}
	if root.Left !=nil&&root.Left.Left==nil&&root.Left.Right==nil{
		sum += root.Left.Val
	}
	sum = r(root.Left,sum)
	sum = r(root.Right,sum)
	return sum
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



