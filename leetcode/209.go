package leetcode

import "math"

func MinSubArrayLen(target int, nums []int) int {

	/**
	双指针 滑动窗口
	本题中起始指针无法向前移动 利用这一点来解题,先找到终止位置,再向后移动起始指针
	每次比较滑动窗口的长度 取最小值
	*/
	start := 0     // 起始指针
	sum := 0       // 求和 临时值
	subLength := 0 // 滑动窗口大小
	result := math.MaxInt64
	for end := 0; end < len(nums); end++ {
		sum += nums[end]
		for sum >= target {
			sum = sum - nums[start] // 减去起始指针的值,因为起始指针要向后滑动,当前这个值就没用了
			subLength = (end - start) + 1
			if result > subLength {
				result = subLength
			}
			start++ // 起始指针向后移动一位
		}
	}

	if result == math.MaxInt64 {
		return 0
	}
	return result
}
