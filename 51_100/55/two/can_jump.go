package two

func CanJump(nums []int) bool {
	pm := 0 // 区间内最大值
	im := 0 // 当前能到的最大位置
	for i := 0; i < len(nums)-1; i++ {
		// 最远处更新
		if nums[i]+i > pm {
			pm = nums[i] + i
		}
		// 是否可以到结尾
		if pm >= len(nums)-1 {
			return true
		}
		// 是否已经到区间最后
		if i < im {
			continue
		}
		// 无法前进一步
		if pm == i {
			return false
		}
		// 更新区间
		im = pm
	}
	if len(nums) <= 1 {
		return true
	}
	return false
}
