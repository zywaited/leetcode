package four

// 类似快速排序
// 但是不用排序所有，只需要排出K个最大就行
func FindKthLargest(nums []int, k int) int {
	return quickSort(nums, 0, len(nums)-1, k)
}

func quickSort(nums []int, l, r, k int) int {
	// 从大到小排序
	mid := nums[r]
	begin := l
	end := r
	for l < r {
		for l < r && nums[l] >= mid {
			l++
		}
		for l < r && nums[r] <= mid {
			r--
		}
		if l != r {
			nums[l], nums[r] = nums[r], nums[l]
		}
	}
	// 中间值
	if l != end {
		nums[l], nums[end] = nums[end], nums[l]
	}
	if l == k-1 {
		return nums[r]
	}
	if l > k-1 {
		// 左边
		return quickSort(nums, begin, l-1, k)
	}
	// 右边
	return quickSort(nums, l+1, end, k)
}
