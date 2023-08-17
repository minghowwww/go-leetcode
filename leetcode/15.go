package leetcode

import "sort"

/*
*
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，
同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。
*/
func ThreeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	if nums[0] > 0 {
		return [][]int{}
	}

	for i := 0; i < len(nums)-2; i++ {
		right := len(nums) - 1
		left := i + 1
		// 比较是否和移动之前的 元素相等
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for right > left {
			sum := nums[i] + nums[left] + nums[right]
			if sum > 0 {
				right--
			} else if sum < 0 {
				left++
			} else {
				// 找到解
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				// 先移动 然后比较是否和移动之前的 元素相等
				for right > left && nums[right] == nums[right+1] {
					right--
				}
				for right > left && nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}

	return result
}
