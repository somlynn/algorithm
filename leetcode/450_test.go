package leetcode

import (
	"fmt"
	"testing"
)

func TestDeleteNode2(t *testing.T) {
	s := "[3,2,5,null,null,4,10,null,null,8,15,7]"
	c := Constructor6()
	root := c.deserialize(s)
	deleteNode2(root, 5)
	fmt.Println(c.serialize(root))
}
