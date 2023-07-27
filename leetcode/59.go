package leetcode

func GenerateMatrix(n int) [][]int {
	// 使用变量指定二维数组长度
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}

	left := 0
	right := n - 1
	top := 0
	bottom := n - 1
	num := 1
	for num <= n*n {
		// 左到右
		for i := left; i <= right; i++ {
			result[top][i] = num
			num++
		}
		top++
		// 上到下
		for i := top; i <= bottom; i++ {
			result[i][right] = num
			num++
		}
		right--
		// 右到左
		for i := right; i >= left; i-- {
			result[bottom][i] = num
			num++
		}
		bottom--
		// 下到上
		for i := bottom; i >= top; i-- {
			result[i][left] = num
			num++
		}
		left++
	}
	return result
}
