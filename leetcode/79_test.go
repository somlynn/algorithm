package leetcode

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	board := [][]byte{
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
		{'A', 'A', 'A', 'A', 'A', 'A'},
	}
	exist := exist(board, "AAAAAAAAAAAAAAB")
	fmt.Println(exist)
}
