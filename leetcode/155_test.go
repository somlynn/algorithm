package leetcode

import (
	"fmt"
	"testing"
)

func TestGetMin(t *testing.T) {
	stack := Constructor3()
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Pop()
	fmt.Println(stack.GetMin())
}
