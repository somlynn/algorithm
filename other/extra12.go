package other

import (
	"container/heap"
	"math"
)

// K 个有序数组的中位数

// 最小堆
// 用最小堆，每次舍弃最小的，当舍弃的个数达到一半时，
// 如果总个数是奇书，则最后舍弃的那个数就是中位数
// 如果总个数是偶数，则(最后舍弃的那个数+下一个舍弃的数)/2就是中位数
func midNumKSortArray(nums [][]int) float64 {
	temp = nums
	if len(nums) == 0 {
		return 0
	}
	minHeap := &PosHeap{}
	heap.Init(minHeap)
	size := 0
	for i := 0; i < len(nums); i++ {
		if len(nums) == 0 {
			continue
		}
		size += len(nums[i])
		heap.Push(minHeap, Pos{i, 0})
	}
	if size == 0 {
		return 0
	}
	median, cnt := 0, 0
	halfCnt := size / 2
	if size%2 == 1 {
		halfCnt = size/2 + 1
	}
	for cnt <= halfCnt {
		top := heap.Pop(minHeap).(Pos)
		median = nums[top.i][top.j]
		cnt++
		if top.j+1 < len(nums[top.i]) {
			top.j++
			heap.Push(minHeap, top)
		}
	}
	if size%2 == 1 {
		return float64(median)
	}
	top := heap.Pop(minHeap).(Pos)
	return (float64(median) + float64(nums[top.i][top.j])) / 2.0
}

var temp [][]int

type Pos struct {
	i, j int
}

type PosHeap []Pos

func (h PosHeap) Len() int {
	return len(h)
}
func (h PosHeap) Less(i, j int) bool {
	return temp[h[i].i][h[i].j] < temp[h[j].i][h[j].j]
}

func (h PosHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PosHeap) Push(x interface{}) {
	*h = append(*h, x.(Pos))
}

func (h *PosHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// 二分法，类似两个有序数组的中位数

//
func findMedian(nums [][]int) float64 {
	total := getTotalCount(nums)
	if total == 0 {
		return 0
	}
	if total%2 == 1 {
		return float64(findKth(nums, total/2+1))
	}
	return float64(findKth(nums, total/2)+findKth(nums, total/2+1)) / 2.0
}

func findKth(nums [][]int, k int) int {
	start, end := 0, math.MaxInt32
	for start+1 < end {
		mid := start + (end-start)/2
		//  大于mid的个数 >= k 说明 mid小了
		if getAllGTE(nums, mid) >= k {
			start = mid
		} else {
			end = mid
		}
	}
	if getAllGTE(nums, start) >= k {
		return end
	}
	return start
}

func getTotalCount(nums [][]int) int {
	n := 0
	for _, num := range nums {
		n += len(num)
	}
	return n
}

// 获取nums中所有大于等于target的个数
func getAllGTE(nums [][]int, target int) int {
	count := 0
	for _, num := range nums {
		count += getGTE(num, target)
	}
	return count
}

// 获取nums中大于等于target的个数
func getGTE(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}
	start, end := 0, len(nums)-1
	for start+1 < end {
		mid := start + (end-start)/2
		if nums[mid] >= target {
			end = mid
		} else {
			start = mid
		}
	}
	if nums[start] >= target {
		return len(nums) - start
	}
	if nums[end] >= target {
		return len(nums) - end
	}
	return 0
}
