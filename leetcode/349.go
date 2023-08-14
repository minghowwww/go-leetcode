package leetcode

import (
	"sort"
)

/*
*
hash表解法
*/
//func Intersection(nums1 []int, nums2 []int) []int {
//	// 构建结果数据
//	size := 1000
//	tmp := make([]int, size)
//	result := []int{}
//	// 初始值都赋值为1
//	for i := 0; i < len(tmp); i++ {
//		tmp[i] = 0
//	}
//
//	for i := 0; i < len(nums1); i++ {
//		tmp[nums1[i]] = 1
//	}
//
//	for i := 0; i < len(nums2); i++ {
//		if tmp[nums2[i]] == 1 && !contains(result, nums2[i]) {
//			result = append(result, nums2[i])
//		}
//	}
//	return result
//}
//
//func contains(s []int, x int) bool {
//	for _, v := range s {
//		if v == x {
//			return true
//		}
//	}
//	return false
//}

/*
*
双指针解法
*/
func Intersection(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	i := 0
	j := 0
	pre := -1
	result := []int{}
	for i < len(nums1) && j < len(nums2) {
		p1 := nums1[i]
		p2 := nums2[j]
		if p1 == p2 {
			if p1 != pre {
				result = append(result, p1)
				pre = p1
				i++
				j++
			} else {
				j++
			}
		}

		if p2 < p1 {
			j++
		}
		if p2 > p1 {
			i++
		}

	}
	return result
}
