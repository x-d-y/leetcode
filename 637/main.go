package main

import (
	"strconv"
)



//func main()  {
//	root := CreateNodeTree([]string{"1","2","3"})
//	fmt.Println(averageOfLevels(root))
//}

func averageOfLevels(root *TreeNode) (ga []float64) {
	var curNodes []*TreeNode
	curNodes = append(curNodes,root)
	var nextNode []*TreeNode
	var sum float64
	for len(curNodes)>0 {
		for _, v := range curNodes {
			sum += float64(v.Val)
			if v.Left != nil {
				nextNode = append(nextNode, v.Left)
			}
			if v.Right != nil {
				nextNode = append(nextNode, v.Right)
			}
		}
		ga = append(ga, sum/float64(len(curNodes)))
		curNodes = nextNode
		nextNode = nil
		sum = 0.0
	}
	return ga
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

