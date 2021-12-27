package two

func CombinationSum4(nums []int, target int) int {
	// 自顶而下
	mp := make(map[int]int, target+1)
	mp[0] = 1
	var dp func(int) int
	dp = func(tn int) int {
		var (
			count int
			ok    bool
		)
		if count, ok = mp[tn]; ok {
			return count
		}
		for _, num := range nums {
			if tn >= num {
				count = dp(tn - num)
				mp[tn-num] = count
				mp[tn] += count
			}
		}
		return mp[tn]
	}
	dp(target)
	return mp[target]
}
