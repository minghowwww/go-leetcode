package leetcode

/*
*
数组移除元素
*/
func RemoveElement(nums []int, val int) int {
	fast := 0
	low := 0
	for _, num := range nums {
		if num != val {
			nums[low] = nums[fast]
			low++
			fast++
		} else {
			fast++
		}
	}
	return low
}

func RemoveElement1(nums []int, val int) int {
	low := 0
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[low] = nums[fast]
			low++
		}
	}
	return low
}
