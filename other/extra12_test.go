package other

import (
	"fmt"
	"testing"
)

// 1,2,3,4,5,5,7,7,8,9,9,11
func TestMidNumKSortArray(t *testing.T) {
	nums := [][]int{
		{2, 5, 7, 9},
		{1, 4, 5},
		{3, 7, 8, 9, 11},
	}
	fmt.Println(midNumKSortArray(nums))
	fmt.Println(findMedian(nums))
}
