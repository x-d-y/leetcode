package main




func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head==nil {
		return head
	}

	last := head
	fast := head
	slow := head
	index := 0
	for {
		fast = fast.Next

		if fast ==nil{
			break
		}
		if index>=n-1{
			last = slow
			slow = slow.Next
		}
		if fast.Next == nil{
			break
		}

		index++
	}
	if head == slow{
		return head.Next
	}

	last.Next = slow.Next
	return head
}

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



