package leetcode

/*
	最长有效括号
*/

// 动态规划
// dp[i]表示以下标i为结尾的字符串的最长有效括号的长度
// 从左往右遍历字符串,每两个字符检查一次
// 如果s[i] == ')' s[i-1] == '(',则字符串形如 xxxxxx() 因此dp[i]=dp[i-2] + 2

// 如果s[i] == ')' s[i-1] == ')'，则字符串形如 xxxxx))
// 		如果s[i-1]是有效子串的一部分(这个有效子串记为subStr)，如果是s[i]是一个更长的子串的一部分,那么一定对应一个'('，
//		并且它在subStr的前面，即[0,i-1-dp[i-1]]区间内。
// 		如果正好 s[i-dp[i-1]-1] == '(' 这字符串形如 xxk(subStr)
//		那么dp[i] = dp[i-1] + dp[k] + 2，其实k=(i-1)-dp[i-1]-1   即dp[i] = dp[i-1]+dp[i-dp[i-1]-2]+2

// 为了方便下标的判断，dp的长度为 len(s)+2，
// dp[i+2] 为结果
// 时间复杂度: O(N) 空间复杂度: O(N)
func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	dp := make([]int, len(s)+2)
	maxLen := 0
	for i := 1; i < len(s); i++ {
		if s[i] == '(' {
			continue
		}
		if s[i-1] == '(' {
			dp[i+2] = dp[i] + 2

		} else if i-dp[i+1] > 0 && s[i-dp[i+1]-1] == '(' {
			dp[i+2] = dp[i+1] + dp[i-dp[i+1]] + 2
		}
		maxLen = Max(maxLen, dp[i+2])
	}
	return maxLen
}

// 栈
// 对于括号相关的题目，我首先想到的是使用栈，
// 遇到每个(,我们把下标放入栈中，
// 遇到每一个),我们先弹出栈顶元素匹配当前的）
// 1、如果栈为空，说明当前的）没有匹配的(，我们将其下标加入栈中，更新（最后一个没有被匹配的右括号的下标）
// 2、如果栈不为空，当前i-stack[len(stack)-1]是以当前）为结尾的最长有效括号的长度
// 然后比较并更新全局最长有效括号的长度maxLen的值。

// 注意：这里有一个小技巧，一开始栈放入一个虚值，
// 如果一开始栈为空，第一个字符为左括号的时候我们会将其放入栈中，这样就不满足提及的「最后一个没有被匹配的右括号的下标」，
// 为了保持统一，我们在一开始的时候往栈中放入一个值为 -1 的元素。
// 比如 ()xxxxx，如果我们不放入-1,需要特殊处理计算maxLen，为了统一，放入虚值避免了特殊处理。
// 时间复杂度: O(N) 空间复杂度: O(N)
func longestValidParentheses2(s string) int {
	maxLen := 0
	stack := make([]int, 0)
	stack = append(stack, -1)
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLen = Max(maxLen, i-stack[len(stack)-1])
			}
		}
	}
	return maxLen
}

// 两次遍历字符串，第一次从左向右遍历，第二次从右向左遍历
// 不需要额外的空间
// 时间复杂度: O(N) 空间复杂度: O(1)

// left记录(的个数，right记录)的个数
// 当left==right相等时，认为匹配，此时以当前i结尾的字符串的最长有效括号为2*right
// 这里有2个特殊情况：))((,)()
// 这里有个小技巧，从左往右遍历时，if right > left left和right置为0

// 同理从右往左遍历，if left > right left和right置为0
func longestValidParentheses3(s string) int {
	left, right, maxLen := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLen = Max(maxLen, 2*right)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			maxLen = Max(maxLen, 2*left)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLen
}
