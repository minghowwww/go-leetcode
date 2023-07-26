package leetcode

// 二分查找 /**
func Search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		middle := (left + right) / 2
		if nums[middle] < target {
			left = middle + 1
		} else if nums[middle] > target {
			right = middle - 1
		} else {
			return middle
		}
	}
	return -1
}
