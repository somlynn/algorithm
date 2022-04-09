package other

import (
	"fmt"
	"testing"
)

func TestHaveCircularDependency(t *testing.T) {
	fmt.Println(haveCircularDependency([][]int{{0, 2}, {1, 2}, {2, 3}, {2, 4}}, 5))
}
