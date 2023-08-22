package leetcode

import "strings"

func ReverseWords(s string) string {

	words := strings.Fields(s)
	l := 0
	r := len(words) - 1
	for l < r {
		words[l], words[r] = words[r], words[l]
		l++
		r--
	}
	return strings.Join(words, " ")
}

//func reverseWords(s string) string {
//	var result string
//	var word string
//	for i := len(s) - 1; i >= 0; i-- {
//		if s[i] == ' ' {
//			if len(word) > 0 {
//				result += word + " "
//				word = ""
//			}
//		} else {
//			word = string(s[i]) + word
//		}
//	}
//	if len(word) > 0 {
//		result += word
//	}
//	if len(result) > 0 && result[len(result)-1] == ' ' {
//		result = result[:len(result)-1]
//	}
//	return result
//}
