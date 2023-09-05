package getOffer

/*
*
选取左下角为起始搜索位置
*/
func FindNumberIn2DArray(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	i, j := len(matrix)-1, 0
	for i >= 0 && j >= 0 && i < len(matrix) && j < len(matrix[i]) {
		flag := matrix[i][j]
		if flag > target {
			i--
		}
		if flag < target {
			j++
		}
		if flag == target {
			return true
		}
	}
	return false
}
