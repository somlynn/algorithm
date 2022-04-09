package leetcode

import "strings"

/*
	验证IP地址
*/

// 思想：IPv4 和 IPv6 地址均是由特定的分界符隔开的字符串组成，并且每个子字符串具有相同格式。
// 可以将地址分为多个块，然后逐块进行验证.仅当每个块都有效时，该地址才有效。这种方法称为 分治法
// 对于 IPv4 地址，通过界定符 . 将地址分为四块；对于 IPv6 地址，通过界定符 : 将地址分为八块。
// 对于 IPv4 地址的每一块，检查它们是否在 0 - 255 内，且没有前置零。
// 对于 IPv6 地址的每一块，检查其长度是否为 1 - 4 位的十六进制数。

const (
	IPv4    = "IPv4"
	IPv6    = "IPv6"
	Neither = "Neither"
)

func validIPAddress(IP string) string {
	is1 := strings.Contains(IP, ".")
	is2 := strings.Contains(IP, ":")
	if is1 && is2 {
		return Neither
	}
	if is1 {
		return validIPv4(IP)
	}
	if is2 {
		return validIPv6(IP)
	}
	return Neither
}

func validIPv4(IP string) string {
	strs := strings.Split(IP, ".")
	if len(strs) != 4 {
		return Neither
	}
	for _, s := range strs {
		if len(s) == 0 || len(s) > 3 {
			return Neither
		}
		if s[0] == '0' && len(s) > 1 {
			return Neither
		}
		if s > "255" {
			return Neither
		}
	}
	return IPv4
}

func validIPv6(IP string) string {
	strs := strings.Split(IP, ":")
	if len(strs) != 8 {
		return Neither
	}
	hex := "0123456789abcdefABCDEF"
	hexMap := make(map[byte]bool, 22)
	for i := range hex {
		hexMap[hex[i]] = true
	}
	for _, s := range strs {
		if len(s) == 0 || len(s) > 4 {
			return Neither
		}
		for i := range s {
			if !hexMap[s[i]] {
				return Neither
			}
		}
	}
	return IPv6
}
