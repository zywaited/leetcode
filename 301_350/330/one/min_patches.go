package one

func MinPatches(nums []int, n int) int {
	// 贪心就是缺啥数字就补啥数字就行，如果在区间就跳过
	add := 0  // 需要添加的数字个数
	miss := 1 // 区间初始化为空
	index := 0
	for miss <= n {
		if index < len(nums) && miss >= nums[index] {
			// 这种情况只要加上源数字就能满足区间，所以不需要新增数字
			// [1,2,2]当变成区间[1,3]时，miss等于3，加上最后一个2区间变为[1,5]，miss变为6
			miss += nums[index]
			index++
			continue
		}
		// 补当前数字
		// [1,3]
		// 拆分成区间: [1,1] [3,3]
		// 也就是缺2就补2，更新区间[1, 3]，miss变为4
		miss += miss
		add++
	}
	return add
}
