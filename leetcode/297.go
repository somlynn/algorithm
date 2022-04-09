package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

/*
	二叉树的序列化与反序列化
*/

type Codec struct {
}

func Constructor6() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	res := make([]string, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				res = append(res, "null")
			} else {
				res = append(res, fmt.Sprintf("%d", node.Val))
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
			}
			size--
		}
	}
	s := strings.Join(res, ",")

	return fmt.Sprintf("[%s]", s)
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if len(data) == 0 {
		return nil
	}
	strs := strings.Split(data[1:len(data)-1], ",")
	val, _ := strconv.ParseInt(strs[0], 10, 64)
	root := &TreeNode{Val: int(val)}
	queue := make([]*TreeNode, 0)
	parent := root
	isLeftChild := true
	for i := 1; i < len(strs); i++ {
		node := getTreeNode(strs[i])
		if isLeftChild {
			parent.Left = node
		} else {
			parent.Right = node
		}
		if node != nil {
			queue = append(queue, node)
		}
		isLeftChild = !isLeftChild
		if isLeftChild && len(queue) > 0 {
			parent = queue[0]
			queue = queue[1:]
		}
	}
	return root
}

func getTreeNode(s string) *TreeNode {
	if s == "null" {
		return nil
	}
	val, _ := strconv.ParseInt(s, 10, 64)
	return &TreeNode{Val: int(val)}
}
