package leetcode

/*
*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

注意：若 s 和 t 中每个字符出现的次数都相同，则称 s 和 t 互为字母异位词。
*/
func isAnagram(s string, t string) bool {
	hash := [26]int{}
	// 初始化hash表
	for i := 0; i < len(hash); i++ {
		hash[i] = 0
	}
	// s 存储 往里加
	for i := 0; i < len(s); i++ {
		index := byte(s[i]) - byte('a')
		hash[index]++
	}
	// t 存储 往下减
	for i := 0; i < len(t); i++ {
		index := byte(t[i]) - byte('a')
		hash[index]--
	}
	// 判断hash表中是否都为0
	for i := 0; i < len(hash); i++ {
		if hash[i] != 0 {
			return false
		}
	}
	return true
}
