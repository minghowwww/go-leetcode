package leetcode

func ReverseStr(s string, k int) string {
	str := []byte(s)
	for i := 0; i < len(str); i += 2 * k {

		if i+k <= len(str) {
			r(str, i, i+k-1)
		} else {
			r(str, i, len(str)-1)
		}
	}
	return string(str)
}

func r(s []byte, start, end int) {
	left := start
	right := end
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
