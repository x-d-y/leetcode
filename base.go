package leetcode

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

