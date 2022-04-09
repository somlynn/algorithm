package leetcode

/*

 */

// 搜索二维矩阵
// 这题与240题虽然条件不一样，但是符合240题，可以用240题的解法去做 但是两者的时间复杂度不一样
// 240题的解法 时间复杂度为 O(m+n)
func searchMatrix2(matrix [][]int, target int) bool {
	r, c := len(matrix), len(matrix[0])
	i, j := 0, c-1
	for i < r && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else {
			j--
		}
	}
	return false
}

// 二分法
// 时间复杂度为 O(logM*N)
// 若将矩阵每一行拼接在上一行的末尾，则会得到一个升序数组，我们可以在该数组上二分找到目标元素。
func searchMatrix3(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	l, h := 0, m*n-1
	for l <= h {
		mid := l + (h-l)/2
		if x := matrix[mid/n][mid%n]; target == x {
			return true
		} else if target > x {
			l = mid + 1
		} else if target < x {
			h = mid - 1
		}
	}
	return false
}
