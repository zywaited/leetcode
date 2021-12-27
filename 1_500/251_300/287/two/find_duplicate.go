package two

func FindDuplicate(nums []int) int {
	size := len(nums)
	// 小于Mid的数量，如果这个数量小于Mid，则重复值在另一边，否则就在这个
	count := 0

	// 数字区间 1,n
	// left, right并不是nums的索引，而是1-n中间数
	left, right := 1, size-1
	for left < right {
		// 1-n中间数，这里加1，不然会死循环
		mid := left + (right-left)>>1 + 1

		// 找小于mid的数量,每次都需要遍历所有数字
		count = 0
		for _, num := range nums {
			if num >= mid {
				continue
			}
			count++
		}

		// 前面计算的>=
		if count >= mid {
			// 在左边
			// 因为计算mid有加1行为，mid的值可能会重复
			right = mid - 1
			continue
		}

		// 在右边
		left = mid
	}

	return left
}
