package leetcode

import "container/heap"

/*
	前K个高频元素
*/

// 最小堆
func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	countMap := make(map[int]int, len(nums))
	for _, v := range nums {
		countMap[v]++
	}
	h := &EHeap{}
	for key, val := range countMap {
		heap.Push(h, [2]int{key, val})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, 0, k)
	for i := 0; i < k; i++ {
		kv := heap.Pop(h).([2]int)
		ret = append(ret, kv[0])
	}
	return ret
}

type EHeap [][2]int

func (h EHeap) Len() int {
	return len(h)
}
func (h EHeap) Less(i, j int) bool {
	return h[i][1] < h[j][1]
}
func (h EHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *EHeap) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}
func (h *EHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 桶排序
// 用 countMap记录元素出现的次数，然后把这些次数作为桶的下标，进行排序，
// buckets[i] 桶里面的元素是出现i次的所有元素。
//
type bucket []int

func topKFrequent2(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	countMap := make(map[int]int, len(nums))
	for _, v := range nums {
		countMap[v]++
	}
	buckets := make([]*bucket, len(nums)+1, len(nums)+1)
	for k, v := range countMap {
		b := buckets[v]
		if b == nil {
			b = &bucket{}
		}
		*b = append(*b, k)
		buckets[v] = b
	}
	ret := make([]int, 0, k)
	for i := len(buckets) - 1; i >= 0; i-- {
		b := buckets[i]
		if b != nil {
			ret = append(ret, *b...)
		}
		if len(ret) == k {
			break
		}
	}
	return ret
}
