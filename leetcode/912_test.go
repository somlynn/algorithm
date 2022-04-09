package leetcode

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	fmt.Println(sortArray([]int{1, 3, 2, 7, 5, 8, 6}))
	fmt.Println(sortArray(sortArray2([]int{1, 3, 2, 7, 5, 8, 6})))
	fmt.Println(sortArray(sortArrayTree([]int{1, 3, 3, 4, 3, 3, 5})))
	fmt.Println(heapSort([]int{1, 3, 2, 7, 5, 8, 6}))
}
