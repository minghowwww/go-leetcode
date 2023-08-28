package getOffer

func GetLeastNumbers(arr []int, k int) []int {
	quickSort(arr, k, 0, len(arr)-1)
	return arr
}

func quickSort(arr []int, k, l, r int) {
	if l >= r {
		return
	}
	i, j := l, r
	for i < j {
		for i < j && arr[j] >= arr[l] {
			// 左起分界指，一定要从右侧开始比较，防止左1 是一个小指，右侧不需要移动，直接i==j了 把j 换到左侧了
			j--
		}
		for i < j && arr[i] <= arr[l] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[l] = arr[l], arr[i]
	if k < i {
		// 说明 下一个小的数字 在右边
		quickSort(arr, k, l, i-1)
	}
	if k > i {
		// 左侧有大的数字 需要拿出去
		quickSort(arr, k, i+1, r)
	}
}
