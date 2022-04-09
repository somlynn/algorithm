package leetcode

import (
	"fmt"
	"testing"
)

func TestIsCompleteTree(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 4},
		Right: &TreeNode{Val: 5},
	}
	root.Right = &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 6},
	}
	fmt.Println(isCompleteTree2(root))
}
