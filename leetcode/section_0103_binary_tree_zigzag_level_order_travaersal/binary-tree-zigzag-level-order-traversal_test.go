package main

import (
	"fmt"
	"testing"
)

func TestZigzagLevelOrder(t *testing.T) {
	root := &TreeNode{
		Val:1,
		Left:&TreeNode{
			Val:2,
		},
		Right:&TreeNode{
			Val:3,
		},
	}
	fmt.Println(ZigzagLevelOrder(root))
}
