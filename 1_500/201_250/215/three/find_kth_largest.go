package three

// 堆
func FindKthLargest(nums []int, k int) int {
	// 不全局建立堆,当然也可以
	// 建立K个元素的最小堆(堆顶就是答案)
	var heapify func(int)
	heapify = func(i int) {
		l := (i << 1) + 1
		r := (i << 1) + 2
		// 3者找到最小的
		min := i
		if l < k && nums[l] < nums[min] {
			min = l
		}
		if r < k && nums[r] < nums[min] {
			min = r
		}
		if min != i {
			nums[i], nums[min] = nums[min], nums[i]
			// 重新构建最小堆
			heapify(min)
		}
	}
	// 初始化
	for i := (k >> 1) - 1; i >= 0; i-- {
		heapify(i)
	}
	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[0] = nums[i]
			heapify(0)
		}
	}
	return nums[0]
}
