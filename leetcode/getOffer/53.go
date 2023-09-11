package getOffer

func Search(nums []int, target int) int {
	/**
	查找最近一个比target小的数的索引
	和 最近一个比target大的数的索引
	*/
	return findBoundary(nums, target) - findBoundary(nums, target-1)
}

func findBoundary(num []int, tar int) int {
	i, j := 0, len(num)-1
	for i <= j {
		mid := (i + j) / 2
		if num[mid] <= tar {
			i = mid + 1
		} else {
			j = mid - 1
		}
	}
	return i
}
