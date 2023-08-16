package leetcode

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	count := 0
	hash := map[int]int{}
	for _, a := range nums1 {
		for _, b := range nums2 {
			key := a + b
			if hash[key] == 0 {
				hash[key] = 1
			} else {
				hash[key]++
			}
		}
	}

	for _, c := range nums3 {
		for _, d := range nums4 {
			key := 0 - (c + d)
			if hash[key] != 0 {
				count += hash[key]
			}
		}
	}
	return count
}
