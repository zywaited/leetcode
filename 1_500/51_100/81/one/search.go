package one

func Search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)>>1
		if nums[l] == target || nums[m] == target || nums[r] == target {
			return true
		}
		if m == l {
			l = m + 1
			continue
		}
		if nums[m] >= nums[l] {
			if nums[m] == nums[l] && nums[m] == nums[r] {
				// 这时无法判断在那一边
				return Search(nums[l+1:m], target) || Search(nums[m+1:r], target)
			}
			if nums[m] > target && target > nums[l] {
				// 在左边
				r = m - 1
				continue
			}
			if target < nums[l] || target > nums[m] {
				// 在右边
				l = m + 1
				continue
			}
			break
		}
		if nums[m] < target && target < nums[r] {
			// 在右边
			l = m + 1
			continue
		}
		if target < nums[m] || target > nums[l] {
			// 在左边
			r = m - 1
			continue
		}
		break
	}
	return false
}
