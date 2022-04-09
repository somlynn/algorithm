package leetcode

/*
	寻找两个正序数组的中位数
*/

// 二分查找法
// 整体思路：利用二分查找法寻找第K个数
// 时间复杂度O(log(m+n),空间复杂度O(1)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1) + len(nums2)
	if n%2 == 1 {
		k := n / 2
		return float64(getKthElement(nums1, nums2, k+1))
	} else {
		k1, k2 := n/2-1, n/2
		return float64(getKthElement(nums1, nums2, k1+1)+getKthElement(nums1, nums2, k2+1)) / 2
	}

}

// 寻找第k个数
func getKthElement(nums1, nums2 []int, k int) int {
	var i, j int
	for {
		// 如果nums1遍历完了，第k个数 = nums2[j+剩余的数-1]的位置
		if i == len(nums1) {
			return nums2[j+k-1]
		}
		// 如果nums2遍历完了，第k个数 = nums2[i+剩余的数-1]的位置
		if j == len(nums2) {
			return nums1[i+k-1]
		}
		// 当还有1个数时，取nums1[i],nums2[j]最小的
		if k == 1 {
			return Min(nums1[i], nums2[j])
		}
		// 将K分成一半，分别在nums1和nums2中查找
		mid := k / 2
		// i+mid超出len(nums1),取len(nums1)
		idx1 := Min(i+mid, len(nums1)) - 1
		// j+mid超出len(nums2),取len(nums2)
		idx2 := Min(j+mid, len(nums2)) - 1
		// 说明[i,idx1] 都在前k个数内 然后再寻找第k-(idx1-i+1)个数
		if nums1[idx1] <= nums2[idx2] {
			k -= idx1 - i + 1
			i = idx1 + 1
		} else {
			k -= idx2 - j + 1
			j = idx2 + 1
		}
	}
}

// 划分数组
// 总体思路：划分两个数组，使得 nums1[0,i] ,nums1[0,j] < nums1[i+1,m] , nums1[j+1,n],且 i + j = m+n+1/2，然后我们的目的就时找到合适的i，j
// 如果为偶数，中位数=（左半部最大值+右半部最小值）/2
// 如果为奇数，中位数= 左半部最大值
// 时间复杂度O(log(min(m,n)),空间复杂度O(1)
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays2(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	iMin, iMax := 0, m
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := (m+n+1)/2 - i
		// 说明切分的i小了，应该在[i,iMax]之间
		if j != 0 && i != m && nums2[j-1] > nums1[i] {
			iMin = i + 1
		} else if i != 0 && j != n && nums1[i-1] > nums2[j] { // 说明切分的i大了，应该在[iMin,i]之间
			iMax = i - 1
		} else {
			maxLeft := 0
			// 求左边最大值
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = Max(nums1[i-1], nums2[j-1])
			}
			// 如果为奇数，中位数= 左半部最大值
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}
			// 如果为偶数，中位数=（左半部最大值+右半部最小值）/2
			// 求右边的最小值
			minRight := 0
			if i == m {
				minRight = nums2[j]
			} else if j == n {
				minRight = nums1[i]
			} else {
				minRight = Min(nums1[i], nums2[j])
			}

			return float64(maxLeft+minRight) / 2
		}
	}
	return 0
}
