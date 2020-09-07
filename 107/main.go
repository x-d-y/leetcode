package main

//func main() {
//	t15 := TreeNode{15, nil, nil}
//	t7 := TreeNode{7, nil, nil}
//	t20 := TreeNode{20, &t15, &t7}
//	t9 := TreeNode{9, nil, nil}
//	t3 := TreeNode{3, &t9, &t20}
//	levelOrderBottom(&t3)
//}

var maxLayer int
var res [][]int

func levelOrderBottom(root *TreeNode) [][]int {
	res = [][]int{}
	rec(root, 0)
	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func rec(root *TreeNode, layer int) {
	if root != nil {
		if layer > maxLayer {
			maxLayer = layer
		}
		if len(res) <= layer {
			res = append(res, []int{})
		}
		rec(root.Left, layer+1)
		rec(root.Right, layer+1)
		res[layer] = append(res[layer], root.Val)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
