package ignore

func Jump(nums []int) int {
	// 已经各个点到最后的最小步数
	mp := make(map[int]int)
	mp[len(nums)-1] = 0
	// step可以不用，因为mp为空可以跳过，这里只是为了尽可能减少循环
	step := 1
	for i := len(nums) - 2; i >= 0; i-- {
		// 判断在这个区间是否可以到最后
		// step <= j <= min(len(nums)-1, nums[i])
		j := nums[i]
		if j > len(nums)-1 {
			j = len(nums) - 1
		}
		for ; j >= step; j-- {
			count, ok := mp[i+j]
			// 这个不能到
			if !ok {
				continue
			}
			// 直接到最后
			if count == 0 {
				mp[i] = 1
				break
			}
			// 是否符合
			if mp[i] != 0 && mp[i] <= count+1 {
				continue
			}
			mp[i] = count + 1
		}
		// step如果去掉，这块逻辑也就去掉
		if mp[i] != 0 {
			step = 1
			continue
		}
		step++
	}
	return mp[0]
}
