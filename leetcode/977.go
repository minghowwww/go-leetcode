package leetcode

/*
*
有序数组平方,双指针实现
*/
func SortedSquares(nums []int) []int {
	i := 0
	j := len(nums) - 1
	result := make([]int, len(nums))
	k := j
	for i <= j {
		sqrti := nums[i] * nums[i]
		sqrtj := nums[j] * nums[j]
		if sqrti > sqrtj {
			result[k] = sqrti
			i++
		} else {
			result[k] = sqrtj
			j--
		}
		k--
	}
	return result
}
