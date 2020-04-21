package two

func NumberOfSubarrays(nums []int, k int) int {
	// 优化第一种解法的数量计算
	// 每个奇数索引
	oen := make([]int, 0, len(nums))
	poi := -1
	ans := 0
	for index, num := range nums {
		if num&1 == 1 {
			// 记录数据
			oen = append(oen, index)
		}
		if len(oen) < k {
			continue
		}
		if len(oen) > k && len(oen) > 0 {
			// 去除前面多余的，往后滑动一个
			poi = oen[0]
			oen = oen[1:]
		}
		if len(oen) > 0 {
			// 前面的连续偶数和自身
			ans += oen[0] - poi
		}
	}
	return ans
}
