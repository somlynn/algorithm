package leetcode

import (
	"fmt"
	"testing"
)

func TestWidthOfBinaryTree(t *testing.T) {
	s := "[1,3,2,5,3,null,9]"
	c := Constructor6()
	root := c.deserialize(s)
	fmt.Println(widthOfBinaryTree(root))
}
