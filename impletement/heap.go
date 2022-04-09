package impletement

type Heap struct {
	size int
	nums []int
}

func NewHeap() *Heap {
	return &Heap{
		size: 0,
		nums: make([]int, 1),
	}
}

// 上浮
func (h *Heap) swim(k int) {
	for k > 1 && h.nums[k/2] < h.nums[k] {
		h.nums[k/2], h.nums[k] = h.nums[k], h.nums[k/2]
		k = k / 2
	}
}

// 下沉
func (h *Heap) sink(k int) {
	for 2*k <= h.size {
		j := 2 * k
		if j < h.size && h.nums[j] < h.nums[j+1] {
			j++
		}
		if h.nums[j] < h.nums[k] {
			break
		}
		h.nums[j], h.nums[k] = h.nums[k], h.nums[j]
		k = j
	}
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Insert(v int) {
	h.nums = append(h.nums, v)
	h.size++
	h.swim(h.size)
}

func (h *Heap) Delete() int {
	max := h.nums[1]
	h.nums[h.size], h.nums[1] = h.nums[1], h.nums[h.size]
	h.size--
	h.nums = h.nums[:len(h.nums)-1]
	h.sink(1)
	return max
}

func (h *Heap) Peek() int {
	if h.Size() > 0 {
		return h.nums[1]
	}
	return -1
}
