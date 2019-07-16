package one

func FirstMissingPositive(nums []int) int {
	nl := len(nums)

	// 最小正整数最大也就是nl+1
	for i, num := range nums {
		if num > 0 {
			continue
		}
		nums[i] = nl + 1
	}

	// 1-nl之间的数字作为索引，并把当前索引对应值置为负数代表出现过
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		if num > nl {
			continue
		}
		if nums[num-1] <= 0 {
			continue
		}
		nums[num-1] = -nums[num-1]
	}

	// 找到未出现的最小正整数
	pos := 1
	for _, num := range nums {
		if num > 0 {
			break
		}
		pos++
	}
	return pos
}
