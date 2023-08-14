package leetcode

func twoSum(nums []int, target int) []int {
	hash := map[int]int{} //值：索引

	for i := 0; i < len(nums); i++ {
		tmp := target - nums[i]
		val, ok := hash[tmp]
		if ok {
			return []int{val, i}
		} else {
			hash[nums[i]] = i
		}
	}
	return nil
}
