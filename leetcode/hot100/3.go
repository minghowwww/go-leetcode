package hot100

/*
*
滑动窗口 + 哈希表
tmp 哈希表记录每个字符最后出现的位置
*/
func LengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	l, r := 0, 0
	b := []byte(s)
	res := 0
	tmp := make(map[byte]int)
	for r != len(s) {
		v, ok := tmp[b[r]]
		if !ok {
			if r-l+1 > res {
				res = r - l + 1
			}
		} else {
			if v >= l {
				l = v + 1
			} else {
				if r-l+1 > res {
					res = r - l + 1
				}
			}
		}
		tmp[b[r]] = r
		r++
	}
	return res
}
