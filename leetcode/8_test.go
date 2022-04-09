package leetcode

import (
	"fmt"
	"testing"
)

func TestMyAtoi(t *testing.T) {
	s := "2147483648"
	fmt.Println(myAtoi(s))
}
