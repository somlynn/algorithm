package other

import (
	"fmt"
	"testing"
)

func TestSmallestDifference(t *testing.T) {
	a := []int{4, 8, 1, 9, 6, 4}
	b := []int{6, 2, 0, 5, 8, 7}
	fmt.Println(smallestDifference(a, b))
}
