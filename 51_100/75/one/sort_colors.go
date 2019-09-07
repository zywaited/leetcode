package one

func SortColors(nums []int) {
	zeroIndex := 0            // 0的最右索引
	twoIndex := len(nums) - 1 // 2的最左索引
	// 那么 <= zeroIndex = 0; >= twoIndex = 2; 中间即为1
	index := 0
	for index <= twoIndex {
		if nums[index] == 0 {
			if index != zeroIndex {
				nums[zeroIndex], nums[index] = nums[index], nums[zeroIndex]
			}
			zeroIndex++
			index++ // 从前面换过来一定是1
			continue
		}
		if nums[index] == 2 {
			if index != twoIndex {
				nums[twoIndex], nums[index] = nums[index], nums[twoIndex]
			}
			twoIndex--
			// 这里不+1是因为可能0从后面换过来了
			continue
		}
		index++
	}
}
