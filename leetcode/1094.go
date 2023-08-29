package leetcode

func CarPooling(trips [][]int, capacity int) bool {
	d := [1001]int{}
	for _, ints := range trips {
		d[ints[1]] += ints[0] // 上车点加上人数
		d[ints[2]] -= ints[0] // 下车点减去人数
	}
	sum := 0
	// 统计所有上车的人 是否超过限制阈值
	for _, v := range d {
		sum += v
		if sum > capacity {
			return false
		}
	}
	return true
}
