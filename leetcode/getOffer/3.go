package getOffer

func FindRepeatNumber(nums []int) int {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if hash[nums[i]] == 0 {
			hash[nums[i]] = 1
		} else {
			return nums[i]
		}
	}
	return -1
}

//func findRepeatNumber(nums []int) int {
//
//	sort.Ints(nums)
//	for i := 0; i < len(nums); i++ {
//		if nums[i] == nums[i+1] {
//			return nums[i]
//		}
//	}
//	return -1
//}
