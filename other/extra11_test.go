package other

import (
	"fmt"
	"testing"
)

func TestCutWood(t *testing.T) {
	fmt.Println(cutWood([]int{4, 7, 2, 10, 5}, 5))
	fmt.Println(cutWood2([]int{4, 7, 2, 10, 5}, 5))
}
