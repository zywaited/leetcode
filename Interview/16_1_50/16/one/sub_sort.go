package one

func SubSort(array []int) []int {
	if len(array) <= 1 {
		return []int{-1, -1}
	}
	m := array[0]
	// 有序的最后索引
	l := -1
	r := -1
	for i := 1; i < len(array); i++ {
		if array[i] > m {
			m = array[i]
		}
		if r == -1 && array[i] >= array[i-1] {
			l = i
			continue
		}
		if array[i] >= m {
			continue
		}
		// 这里可以二分，为了简单就循环了
		for ; l >= 0; l-- {
			if array[i] >= array[l] {
				break
			}
		}
		r = i
	}
	l++
	if r <= l {
		return []int{-1, -1}
	}
	return []int{l, r}
}
