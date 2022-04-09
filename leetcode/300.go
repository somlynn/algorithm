package leetcode

/*
	最长递增子序列
*/

// 贪心+二分查找
// 如果我们想让上升子序列尽可能的长，我们需要让序列差值尽可能的小，因此我们希望每次在上升子序列最后加上的那个数尽可能的小。
// tails 表示结果队列，二分法用来找出下一符合的元素插入的位置
// 比如：nums= [10,9,2,5,3,7,21,18]
// 10 idx = 0 tails=[10]
// 9 idx = 0 tails =[9]
// 2 idx = 0 tails = [2]
// 5 idx = 1 tails =[2,5]
// 3 idx = 1 tails = [2,3]
// 7 idx = 2 tails = [2,3,7]
// 21 idx = 3 tails = [2,3,7,21]
// 18 idx = 3 tails = [2,3,7,18]
func lengthOfLIS(nums []int) int {
	ret := make([]int, len(nums))
	size := 0
	for _, num := range nums {
		idx := binarySearch(ret, size, num)
		ret[idx] = num
		// 当下一个插入位置是 size时，即idx == size，说明插入新元素，size++，idx < size 是更新元素。
		if idx == size {
			size++
		}
	}
	return size
}

// 二分法找出 仅此key的下一个位置，
// 所以 tail[mid] > key 时，key的插入位置可能是[l,m] 所以下一个搜索范围为[l,m] h=m ,
// 这里解释一下为啥不是h =m-1,因为 tails[m] > key ，则可能 tails[m-1] <= key ,那么key的插入位置是 m ,所以下一个搜索范围为[l,m]
// 当 tail[mid] <= key 是，key的插入位置可能是(m,h]，所以下一个搜索范围为[m+1,h] l=m+1
func binarySearch(ret []int, size, key int) int {
	l, h := 0, size
	for l < h {
		m := l + (h-l)/2
		if key == ret[m] {
			return m
		} else if key < ret[m] {
			h = m
		} else {
			l = m + 1
		}
	}
	return l
}

// 动态规划
// dp[i]表示[0,i]的最长上升子序列
// dp[i] = max(dp[j])+1 并且0<=j <i && nums[j] < nums[i]

// 判断dp[i-1]中加入nums[i]形成dp[i],
// 首先需要 从 [0,i-1] 中找出 比nums[i] 小的数来组成递增子序列。 比如此时找到了，m,n,k三个
// 然后在这三个上拼上nums[i] 即 dp[m]+1,dp[n]+1,dp[k]+1 取最大的
// 即 max(dp[m],dp[n],dp[k])+1 ，
// maxSize 记录整个数组最大子序列的长度，maxLen 记录 max(max(dp[m],dp[n],dp[k])).
// 时间复杂度: O(N^2) 空间复杂度: O(N)
func lengthOfLIS2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	maxSize := 1
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		maxLen := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				maxLen = Max(maxLen, dp[j])
			}
		}
		dp[i] = maxLen + 1
		maxSize = Max(maxSize, dp[i])
	}
	return maxSize
}
