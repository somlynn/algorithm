package leetcode

/*
	Z字形变换
*/

type Bucket []byte

func convert(s string, numRows int) string {
	if len(s) == 0 || numRows == 1 {
		return s
	}
	buckets := make([]Bucket, numRows)
	cur := 0
	flag := true
	for i := range s {
		if cur == 0 {
			flag = true
		}
		if cur == numRows-1 {
			flag = false
		}
		if buckets[cur] == nil {
			buckets[cur] = Bucket{s[i]}
		} else {
			buckets[cur] = append(buckets[cur], s[i])
		}
		if flag {
			cur++
		} else {
			cur--
		}
	}
	ret := make([]byte, 0, len(s))
	for _, v := range buckets {
		ret = append(ret, v...)
	}
	return string(ret)
}
