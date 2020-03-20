package one

func GetLeastNumbers(arr []int, k int) []int {
	// 快排逻辑（拆分区域，这样可以不用全排序）
	qs(arr, 0, len(arr)-1, k)
	return arr[:k]
}

// 返回分离点
func qs(arr []int, start, end, k int) int {
	if start >= end {
		return start
	}
	base := arr[start]
	low := start
	high := end
	for low < high {
		for low < high && arr[high] >= base {
			high--
		}
		arr[low] = arr[high]
		for low < high && arr[low] <= base {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = base
	if low == k {
		return low
	}
	if low > k-1 {
		// 左边过多
		return qs(arr, 0, low-1, k)
	}
	return qs(arr, low+1, end, k)
}
