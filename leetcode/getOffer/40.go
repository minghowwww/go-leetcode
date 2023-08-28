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
			j--
		}
		for i < j && arr[i] <= arr[l] {
			i++
		}
		arr[i], arr[j] = arr[j], arr[i]
	}
	arr[i], arr[l] = arr[l], arr[i]
	if k < i {
		quickSort(arr, k, l, i-1)
	}
	if k > i {
		quickSort(arr, k, i+1, r)
	}
}
