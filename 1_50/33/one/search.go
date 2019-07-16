package one

func Search(nums []int, target int) int {
	left := 0
	n := len(nums)
	right := n - 1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		}

		// 如果在左边
		if nums[mid] >= nums[left] {
			// 判断target的位置
			if nums[mid] > target && target >= nums[left] {
				right = mid - 1
				continue
			}
			left = mid + 1
			continue
		}
		// 如果在右边
		if nums[mid] < target && target <= nums[right] {
			left = mid + 1
			continue
		}
		right = mid - 1
	}
	return -1
}
