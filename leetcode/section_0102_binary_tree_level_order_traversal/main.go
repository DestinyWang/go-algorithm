package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	rst := make([][]int, 0)
	if root == nil {
		return rst
	}
	currLevel := 1
	q := make([]*TreeNode, 0)
	level := make([]int, 0)
	// pop
	q = append(q, root)
	level = append(level, currLevel)
	buf := make([]int, 0)
	for i := 0; i < len(q); i++ {
		// push
		t := q[i]
		if t.Left != nil {
			q = append(q, t.Left)
			level = append(level, level[i]+1)
		}
		if t.Right != nil {
			q = append(q, t.Right)
			level = append(level, level[i]+1)
		}
		if currLevel != level[i] {
			rst = append(rst, buf)
			buf = make([]int, 0)
			currLevel++
		}
		buf = append(buf, q[i].Val)
	}
	if len(buf) > 0 {
		rst = append(rst, buf)
	}
	return rst
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:3,
			},
			Right:&TreeNode{
				Val:4,
				Left:&TreeNode{
					Val:5,
				},
			},
		},
	}
	order := levelOrder(root)
	fmt.Println(order)
}
