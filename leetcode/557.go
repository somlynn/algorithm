package leetcode

func reverseWords2(s string) string {
	i := 0
	bs := []byte(s)
	for i < len(bs) {
		start := i
		for i < len(bs) && bs[i] != ' ' {
			i++
		}
		end := i - 1
		for start < end {
			bs[start], bs[end] = bs[end], bs[start]
			start++
			end--
		}
		for i < len(bs) && i == ' ' {
			i++
		}
	}
	return string(bs)
}
