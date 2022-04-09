package leetcode

/*
	排序数组
*/

// 快速排序
// 时间复杂度: O(NlogN) 空间复杂度O(logN)
// 如果数组是有序的，快速排序会退化成 O(N^2),
// 如果数组存在大量重复元素，快速排序也会退化成O(N^2)

// 优化方法：
// 对于小数组切换成插入排序，
// 数组尽量随机化
// 对于大量重复的元素，使用三向切分快速排序
func sortArray(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

// 三向切分
func sortArrayTree(nums []int) []int {
	quickSort2(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, l, h int) {
	if l >= h {
		return
	}
	j := partition(nums, l, h)
	quickSort(nums, l, j-1)
	quickSort(nums, j+1, h)
}

func partition(nums []int, l, h int) int {
	v := nums[l]
	i, j := l+1, h
	for {
		for ; i < h && nums[i] <= v; i++ {
		}
		for ; j > l && nums[j] >= v; j-- {
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}

// 三向切分
func quickSort2(nums []int, l, h int) {
	if l >= h {
		return
	}
	base := nums[l]
	lt, eq, gt := l, l, h
	for eq <= gt {
		if nums[eq] > base {
			nums[eq], nums[gt] = nums[gt], nums[eq]
			gt--
		} else if nums[eq] < base {
			nums[lt], nums[eq] = nums[eq], nums[lt]
			lt++
			eq++
		} else {
			eq++
		}
	}
	quickSort2(nums, l, lt-1)
	quickSort2(nums, gt+1, h)
}

var aux []int

// 归并排序
// 是稳定的排序，时间复杂度:O(NlogN),空间复杂度O(N)
func sortArray2(nums []int) []int {
	aux = make([]int, len(nums))
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, l, h int) {
	if l >= h {
		return
	}
	m := l + (h-l)/2
	mergeSort(nums, l, m)
	mergeSort(nums, m+1, h)
	merge4(nums, l, m, h)
}

func merge4(nums []int, l, m, h int) {
	i, j := l, m+1
	for k := l; k <= h; k++ {
		aux[k] = nums[k]
	}
	for k := l; k <= h; k++ {
		if i > m {
			nums[k] = aux[j]
			j++
		} else if j > h {
			nums[k] = aux[i]
			i++
		} else if aux[i] < aux[j] {
			nums[k] = aux[i]
			i++
		} else {
			nums[k] = aux[j]
			j++
		}
	}
}

// 插入排序
func sortArray3(nums []int) []int {

	return nil
}

func heapSort(nums []int) []int {
	size := len(nums)
	for k := size/2 - 1; k >= 0; k-- {
		sink(nums, k, size)
	}
	for size > 1 {
		nums[0], nums[size-1] = nums[size-1], nums[0]
		size--
		sink(nums, 0, size)
	}
	return nums
}

func sink(nums []int, k, size int) {
	for 2*k+1 < size {
		j := 2*k + 1
		if j < size-1 && nums[j] < nums[j+1] {
			j++
		}
		if nums[k] > nums[j] {
			break
		}
		nums[k], nums[j] = nums[j], nums[k]
		k = j
	}
}

// 外部排序
// 能够处理极大量数据的排序算法。通常来说，外排序处理的数据不能一次装入内存，只能放在读写较慢的外存储器（通常是硬盘）上。外排序通常采用的是一种“排序-归并”的策略。
// 在排序阶段，先读入能放在内存中的数据量，将其排序输出到一个临时文件，依此进行，将待排序数据组织为多个有序的临时文件。而后在归并阶段将这些临时文件组合为一个大的有序文件，也即排序结果。

// 比如有10000记录的文件，内存只能存储1000个记录。我们可以将文件分成10个1000记录的临时文件，然后进行多路归并排序。
// 将排序的结果在存储在一个大文件中，然后在对每一个大文件进行归并排序。
// 注意内存有限，归并的时候不能一次全部加载到内存，只能读取一部分。
// 对于外部排序算法来说，影响整体排序效率的因素主要取决于读写外存的次数，即访问外存的次数越多，算法花费的时间就越多，效率就越低。
// 读写IO的次数与归并次数有关。
