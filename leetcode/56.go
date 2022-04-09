package leetcode

import "sort"

/*
	合并区间
*/

// 排序+两个指针 ，这两个指针表示合并之后的start和end
// 首先按照 区间的第一个元素进行排序
// 这样我们只需要考虑2中情况
// 1、 [2,5] [6,9]
// 2、 [2,6] [4,5]

// 对于第一种情况，if intervals[i][0] > end 我们不需要合并，则将[2,5]加入结果集中，更新 start ，end
// [6,9]在和后面的区间比较
// 对于第二种情况，我们只要更新end =  Max(end,intervals[i][1])
// 不管第一种情况和第二种情况，end 都是要更新的 ，所以这里把 end = Max(end,intervals[i][1])放在外边，因为不需要判断
// 注意：for循环结束后，最后一个合并区间还没有加入到结果集中，所以最后别忘了ret = append(ret,[]int{start,end})
func mergeRange(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}
	ret := make([][]int, 0)
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	start, end := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > end {
			ret = append(ret, []int{start, end})
			start = intervals[i][0]
		}
		end = Max(end, intervals[i][1])
	}
	// 加入最后一个区间
	ret = append(ret, []int{start, end})
	return ret
}
