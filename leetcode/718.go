package leetcode

/*
	最长重复子数组
*/

// 动态规划
// 记dp[i][j] 为 nums1[i:] 和nums2[j:]以i,j为前缀的最长重复子数组
// a:[1,2,3,2,1],b[3,2,1,4,7]
// 当前 nums1[i] == nums[j] 时，dp[i][j]= dp[i+1][j+1] + 1
// 当前 nums1[i] != nums[j] 时，则不存在i,j 为公共前缀的数组，因此dp[i][j] = 0
// 时间复杂度: O(N*M) 空间复杂度: O(N*M)
func findLength(nums1 []int, nums2 []int) int {
	m, n := len(nums1), len(nums2)
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	maxLen := 0
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
			} else {
				dp[i][j] = 0
			}
			maxLen = Max(maxLen, dp[i][j])
		}
	}
	return maxLen
}

// 滑动窗口
// 非常好理解，想象两把尺子，错开之后比较相同的部分，找最长相同的串就好了。
// 这里有一个需要注意，小数组不动，大数组移动。有三种情况：
// 		12321
//	 632147
//  重合部分是 nums1:[0,i] nums2: [n-i] 重合的size=i
//     12321
//   7632147
// 重合部分是 nums1:[0,i] nums2: [j,m] j=[0,n-m] 重合的size = m
//     12321
//     	7632147
// 重合部分是 nums1:[i:m-1) nums2 [0,m-i),重合的size = m-i
// 每次重合求出公共重复长度,对于所有的重合，比较求出maxLen
// 时间复杂度: O((N+M)*min(N,M)) 空间复杂度: O(1)
func findLength2(nums1 []int, nums2 []int) int {
	if len(nums1) > len(nums2) {
		return findLength2(nums2, nums1)
	}
	m, n := len(nums1), len(nums2)
	maxLen := 0
	for i := 1; i <= m; i++ {
		count := maxLength(nums1, nums2, 0, n-i, i)
		maxLen = Max(maxLen, count)
	}
	for j := n - m; j >= 0; j-- {
		count := maxLength(nums1, nums2, 0, j, m)
		maxLen = Max(maxLen, count)
	}
	for i := 1; i < m; i++ {
		count := maxLength(nums1, nums2, i, 0, m-i)
		maxLen = Max(maxLen, count)
	}
	return maxLen

}

// 找出两个数组重合部分最长的公共长度
// i,j 表示nums1和nums2的开始遍历位置,size表示遍历的长度
func maxLength(nums1, nums2 []int, i, j, size int) int {
	count, maxLen := 0, 0
	for k := 0; k < size; k++ {
		if nums1[i+k] == nums2[j+k] {
			count++
		} else {
			count = 0
		}
		maxLen = Max(maxLen, count)
	}
	return maxLen
}
