package other

import (
	"fmt"
	"testing"
)

func TestMoveNums(t *testing.T) {
	nums := []int{2, -1, -2, 7, 8, -3, 9, -4, 5}
	moveNums(nums)
	fmt.Println(nums)
}
