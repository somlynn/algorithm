package other

import (
	"fmt"
	"testing"
)

func TestSortOddEvenList(t *testing.T) {
	head := BuildListNode([]int{1, 8, 3, 6, 5, 4, 7, 2})
	head = sortOddEvenList(head)
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
