package hot100

/*
*
方案1 暴力法 时间复杂度太高 不推荐
*/
//func LongestPalindrome(s string) string {
//	res := ""
//	l, r := 0, len(s)-1
//	for l < len(s) && r >= l {
//		if isEqual(l, r, s) && r-l+1 > len(res) {
//			res = s[l : r+1]
//			l++
//			r = len(s) - 1
//		} else {
//			r--
//		}
//
//		if r == l || len(res) > r-l+1 {
//			l++
//			r = len(s) - 1
//		}
//	}
//	return res
//}
//
//func isEqual(i, j int, s string) bool {
//	if i == j {
//		return true
//	} else {
//		if i == j-1 {
//			return s[i] == s[j]
//		}
//	}
//	if s[i] == s[j] {
//		i++
//		j--
//		return isEqual(i, j, s)
//	}
//	return false
//}

/*
*
方案二 中心扩展
*/
func LongestPalindrome(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		// 一个元素为中心
		l1, r1 := expand(s, i, i)
		// 两个元素为中心
		l2, r2 := expand(s, i, i+1)
		if r1-l1 > end-start {
			start, end = l1, r1
		}
		if r2-l2 > end-start {
			start, end = l2, r2
		}
	}
	return s[start : end+1]
}

/*
*
相等的值向两边扩展
*/
func expand(s string, i, j int) (int, int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	// 回溯
	return i + 1, j - 1
}
