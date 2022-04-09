package other

import (
	"fmt"
	"testing"
)

func TestRangeMaxSum(t *testing.T) {
	nums := []int{6, 2, 1}
	fmt.Println(rangeMaxSum(nums))
	fmt.Println(rangeMaxSum2(nums))
}
