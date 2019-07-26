package two

func Jump(nums []int) int {
	// 尽可能的跳到最远处
	// 注意：不是直接nums[i]
	// 比如: 2, 3, 1, 1，1，不是直接到1，而是先到3这里, 4, 3, 1, 1, 1，才是直接就到最后一个1
	// 也就是最远为：max(i+nums[i])
	// 贪心算法在这里和广度优先算法类似，但减少了重复计算，效率更高

	pm := 0 // 区间内最大值
	im := 0 // 当前能到的最大位置
	step := 0
	for i := 0; i < len(nums)-1; i++ {
		// 最远处更新
		if nums[i]+i > pm {
			pm = nums[i] + i
		}
		// 是否可以到结尾
		if pm >= len(nums)-1 {
			return step + 1
		}
		// 是否已经到区间最后
		if i < im {
			continue
		}
		// 走一步
		step++
		// 更新区间
		im = pm
	}
	return step
}
