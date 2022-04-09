package leetcode

import (
	"fmt"
	"testing"
)

func TestFindDiagonalOrder(t *testing.T) {
	res := findDiagonalOrder([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	fmt.Println(res)
}
