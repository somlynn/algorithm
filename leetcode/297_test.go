package leetcode

import (
	"fmt"
	"testing"
)

func TestSerializeAndDeserialize(t *testing.T) {
	s := "[1,2,3,null,null,4,5,null,null,null,null]"
	c := Constructor6()
	root := c.deserialize(s)
	fmt.Print(c.serialize(root))
}
