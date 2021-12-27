package one

func MoveZeroes(nums []int) {
	// 0的索引
	zeroIndex := -1
	// 直到找到一个不是0的数字
	// 然后与零值交换
	for i, num := range nums {
		if num != 0 && zeroIndex != -1 {
			nums[zeroIndex], nums[i] = nums[i], nums[zeroIndex]
			zeroIndex++ // 下一个0
			continue
		}
		if num == 0 && zeroIndex == -1 { // 找到0
			zeroIndex = i
		}
	}
}
