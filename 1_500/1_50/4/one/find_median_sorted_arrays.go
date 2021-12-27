package one

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	// 左边一定要比右边的小，这样不需要考虑两种情况
	if m > n {
		nums1, nums2 = nums2, nums1
		m, n = n, m
	}

	iMin := 0
	iMax := m
	// 如果有需要需要处理一下边界值，因为m+n+1可能过边界
	half := (m + n + 1) / 2
	odd := (m+n)%2 != 0
	for iMin <= iMax {
		i := (iMin + iMax + 1) / 2
		if i > iMax {
			i = iMax
		}

		j := half - i

		// 过大
		if i > 0 && j < n && nums1[i-1] > nums2[j] {
			iMax -= 1
			continue
		}

		// 过小
		if j > 0 && i < iMax && nums2[j-1] > nums1[i] {
			iMin += 1
			continue
		}

		maxLeft := 0
		hasLeft := false
		if i > 0 {
			maxLeft = nums1[i-1]
			hasLeft = true
		}

		if j > 0 && (!hasLeft || nums2[j-1] > maxLeft) {
			maxLeft = nums2[j-1]
		}

		if odd {
			return float64(maxLeft)
		}

		minRight := 0
		hasRight := false
		if i < m {
			minRight = nums1[i]
			hasRight = true
		}

		if j < n && (!hasRight || nums2[j] < minRight) {
			minRight = nums2[j]
		}

		return float64(maxLeft+minRight) / 2.0
	}

	return 0.0
}
