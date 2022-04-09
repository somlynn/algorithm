package leetcode

import (
	"container/heap"
	"math"
)

/*
	滑动窗口最大值
*/

// 超出时间限制，对于随机的数组执行挺快，但是最坏情况就退化为O(n^2)
func maxSlidingWindow(nums []int, k int) []int {
	size := len(nums)
	if size == 0 {
		return []int{}
	}
	res := make([]int, size-k+1)
	maxIdx, maxVal := -1, math.MinInt32
	for i := 0; i < size-k+1; i++ {
		if maxIdx >= i {
			if nums[i+k-1] > maxVal {
				maxVal = nums[i+k-1]
				maxIdx = i + k - 1
			}
		} else {
			maxVal = nums[i]
			for j := i; j < i+k; j++ {
				if maxVal < nums[j] {
					maxVal = nums[j]
					maxIdx = j
				}
			}
		}
		res[i] = maxVal
	}
	return res
}

// 双向队列、单调递减队列
// 1、从左向右遍历元素
// 2、队列的最右端始终保存当前元素，队列从右向左都是大于当前元素的数的下标，
// 3、队列中最大的数，即队列最左端，与当前元素的距离是否小于等于K，即最大的数是否在窗口内。如果不在，缩小窗口的左端。
// 4、如果在窗口内，并且已经形成了长度为K的窗口，即i+1 >= k 则加入结果集内
func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 1 {
		return nums
	}
	queue := make([]int, 0)
	res := make([]int, 0, len(nums)-k+1)
	for i := 0; i < len(nums); i++ {
		// 当队列不为空，且队尾小于当前元素时，弹出保证队列为递减队列
		for len(queue) != 0 && nums[queue[len(queue)-1]] <= nums[i] {
			queue = queue[:len(queue)-1]
		}
		queue = append(queue, i)
		// 不在滑动窗口内，弹出队首
		if queue[0] <= i-k {
			queue = queue[1:]
		}
		// 当形成滑动窗口时，才加入结果集中
		if i+1 >= k {
			res = append(res, nums[queue[0]])
		}
	}
	return res
}

type Item struct {
	idx int
	val int
}

type maxHeap []Item

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i].val > h[j].val
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(v interface{}) {
	*h = append(*h, v.(Item))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// 优先队列
// 时间复杂度为 O(nlogk)，空间复杂度为O(n)
func maxSlidingWindow3(nums []int, k int) []int {
	h := &maxHeap{}
	// 先初始化
	for i := 0; i < k; i++ {
		*h = append(*h, Item{idx: i, val: nums[i]})
	}
	// 构建堆
	heap.Init(h)
	res := make([]int, 0, len(nums)-k+1)
	// 先将第一个滑动窗口的最大值放进去
	res = append(res, (*h)[0].val)
	for i := k; i < len(nums); i++ {
		heap.Push(h, Item{idx: i, val: nums[i]})
		// 当前最大值不在滑动窗口范围内,则弹出
		for (*h)[0].idx <= i-k {
			heap.Pop(h)
		}
		res = append(res, (*h)[0].val)
	}
	return res
}
