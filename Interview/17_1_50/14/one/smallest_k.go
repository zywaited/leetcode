package one

// 直接采用K个数堆排序
func SmallestK(arr []int, k int) []int {
	nums := arr[:k]
	if k > 0 {
		var heapify func(int)
		heapify = func(i int) {
			l := i<<1 + 1
			r := l + 1
			max := i
			if l < k && nums[l] > nums[max] {
				max = l
			}
			if r < k && nums[r] > nums[max] {
				max = r
			}
			if max != i {
				nums[i], nums[max] = nums[max], nums[i]
				heapify(max)
			}
		}
		// 建立最大堆
		for i := k >> 1; i >= 0; i-- {
			heapify(i)
		}
		for i := k; i < len(arr); i++ {
			if arr[i] < nums[0] {
				nums[0] = arr[i]
				heapify(0)
			}
		}
	}
	return nums
}
