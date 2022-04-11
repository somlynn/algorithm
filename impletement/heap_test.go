package impletement

import (
	"fmt"
	"testing"
)

func TestHeap(t *testing.T) {
	heap := NewHeap()
	heap.Push(3)
	heap.Push(4)
	heap.Push(6)
	heap.Push(5)
	heap.Push(2)
	fmt.Println(heap.Peek())
	fmt.Println(heap.Pop())
	fmt.Println(heap.Peek())
}
