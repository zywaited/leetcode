package one

// 堆排序(大顶堆)
func GetLeastNumbers(arr []int, k int) []int {
	ans := arr[:k]
	for index := k >> 1; index >= 0; index-- {
		heapify(ans, index)
	}
	for index := k; index < len(arr); index++ {
		if len(ans) > 0 && arr[index] < ans[0] {
			ans[0] = arr[index]
			heapify(ans, 0)
		}
	}
	return ans
}

func heapify(arr []int, index int) {
	l := index<<1 + 1
	r := l + 1
	max := index
	if l < len(arr) && arr[l] > arr[max] {
		max = l
	}
	if r < len(arr) && arr[r] > arr[max] {
		max = r
	}
	if max != index {
		arr[index], arr[max] = arr[max], arr[index]
		heapify(arr, max)
	}
}
