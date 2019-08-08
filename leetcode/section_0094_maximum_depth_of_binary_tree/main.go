package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 非递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := make([]*TreeNode, 0)
	d := make([]int, 0)
	// init
	q = append(q, root)
	d = append(d, 1)
	max := 0
	for i := 0; i < len(q); i++ {
		t := q[i]
		currDepth := d[i]
		if t.Left != nil {
			q = append(q, t.Left)
			d = append(d, currDepth+1)
		}
		if t.Right != nil {
			q = append(q, t.Right)
			d = append(d, currDepth+1)
		}
	}
	for i := 0; i < len(d); i++ {
        if max < d[i] {
        	max = d[i]
        }
	}
	return max
}

// 递归
func maxDepth_rev(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		ld := maxDepth_rev(root.Left)
		rd := maxDepth_rev(root.Right)
		return max(ld, rd) + 1
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	root := &TreeNode{
		Val:3,
		Left: &TreeNode{
			Val:9,
		},
		Right: &TreeNode{
			Val:20,
			Left:&TreeNode{
				Val:15,
			},
			Right:&TreeNode{
				Val:7,
			},
		},
	}
	fmt.Println(maxDepth(root))
}
