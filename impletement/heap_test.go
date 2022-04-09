package impletement

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	heap := NewHeap()
	heap.Insert(3)
	heap.Insert(4)
	heap.Insert(6)
	heap.Insert(5)
	heap.Insert(2)
	fmt.Println(heap.Peek())
	fmt.Println(heap.Delete())
	fmt.Println(heap.Peek())
}
