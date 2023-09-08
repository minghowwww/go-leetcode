package getOffer

func Exchange(nums []int) []int {

	l, r := 0, len(nums)-1
	for l < r {
		for l < r && nums[l]&1 == 1 {
			// 从左往右找偶数
			l++
		}
		if l == r {
			return nums
		} else {
			// 从右往左找奇数
			for l < r && nums[r]&1 == 0 {
				r--
			}
			// 交换
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	return nums
}
