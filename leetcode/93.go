package leetcode

import (
	"strconv"
	"strings"
)

/*
	复原IP地址
*/

// 回溯+剪枝
// 25525511135 以它为例，我们可以2,25,225 为第一个片段,然后再切出第二个片段，如此
// 用path保存这些片段, start遍历的开始位置
// 当len(path)最终为4，start == len(s)时，结果正好是IP地址,保存结果
// 当len(path)==4 ，start < len(s)时，说明已经切分4个片段后，还没遍历完数组，不是解，直接return
// IP地址一个片段最多3位 ，从start开始，截取长度size为1,2,3 的片段看看是否满足IP地址的要求，不满足直接跳过（剪枝）

//  1、如果start+size >= len(s) 跳过
//  2、如果 size > 1，切出的长度大于1的片段有前导0 ,跳过
//  3、如果 片段对于的数值大于255，跳过
//  否则，是符合条件的片段，加入path中。第一个片段的一种可能下（a），然后尝试第二个片段 backtrack(start+size)
//  在以a为第一段片段的遍历完后，我们回溯为原来的状态，path = path[:len(path)-1]
//  然后尝试以b作为一段片段的递归
func restoreIpAddresses(s string) []string {
	ret := make([]string, 0)
	path := make([]string, 0)
	var backtrack func(int)
	backtrack = func(i int) {

		if len(path) == 4 && i == len(s) {
			ret = append(ret, strings.Join(path, "."))
			return
		}

		if len(path) == 4 && i < len(s) {
			return
		}
		for size := 1; size <= 3; size++ {
			if i+size-1 >= len(s) {
				return
			}
			// 前导零
			if size > 1 && s[i] == '0' {
				return
			}
			str := s[i : i+size]
			if v, _ := strconv.Atoi(str); v > 255 {
				return
			}
			path = append(path, str)
			backtrack(i + size)
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return ret
}
