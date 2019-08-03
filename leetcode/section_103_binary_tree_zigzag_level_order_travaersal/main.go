package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	rst := make([][]int, 0)
	if root == nil {
		return rst
	}
	q := make([]*TreeNode, 0)
	currLevel := 1
	l := make([]int, 0)
	buf := make([]int, 0)
	// push
	q = append(q, root)
	l = append(l, 1)
	for i := 0; i < len(q); i++ {
		if q[i].Left != nil {
			q = append(q, q[i].Left)
			l = append(l, l[i]+1)
		}
		if q[i].Right != nil {
			q = append(q, q[i].Right)
			l = append(l, l[i]+1)
		}
		if l[i] == currLevel {
			// 如果与当前处理深度一致
			buf = append(buf, q[i].Val)
		} else {
			// if not equals currLevel
			if currLevel % 2 == 0 {
				buf = reverse(buf)
			}
			rst = append(rst, buf)
			buf = make([]int, 0)
			buf = append(buf, q[i].Val)
			currLevel ++
		}
	}
	if len(buf) != 0 {
		// if buf is not empty
		if currLevel % 2 == 0 {
			buf = reverse(buf)
		}
		rst = append(rst, buf)
	}
	return rst
}

func reverse(a1 []int) (a2 []int) {
	a2 = make([]int, 0)
	for i := len(a1) - 1; i >= 0; i-- {
		a2 = append(a2, a1[i])
	}
	return a2
}

func main() {
	root := &TreeNode{
		Val:1,
		Left:&TreeNode{
			Val:2,
		},
		Right:&TreeNode{
			Val:3,
		},
	}
	fmt.Println(zigzagLevelOrder(root))
}
