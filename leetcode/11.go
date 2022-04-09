package leetcode

/*
	盛最多水的容器
*/

// 双指针
func maxArea(height []int) int {
	maxArea, l, r := 0, 0, len(height)-1
	for l < r {
		maxArea = Max(maxArea, Min(height[l], height[r])*(r-l))
		if height[l] <= height[r] {
			l++
		} else {
			r--
		}
	}
	return maxArea
}

func maxArea2(height []int) int {
	maxArea, area, l, r := 0, 0, 0, len(height)-1
	for l < r {
		if hl, hr, dist := height[l], height[r], r-l; hl > hr {
			area = dist * hr
			r--
		} else {
			area = dist * hl
			l++
		}
		if area > maxArea {
			maxArea = area
		}
	}
	return maxArea
}
