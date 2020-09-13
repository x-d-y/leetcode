package leetcode

import "strconv"

func CreateList(list []int)*ListNode{
	var lastNode *ListNode
	lastNode = nil
	var firstNode *ListNode
	for _,v := range list{
		cNode := new(ListNode)
		cNode.Val = v
		cNode.Next = nil
		if lastNode !=nil{
			lastNode.Next = cNode
		}else{
			firstNode = cNode
		}
		lastNode = cNode
	}
	return firstNode
}

type ListNode struct {
	Val int
	Next *ListNode
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


