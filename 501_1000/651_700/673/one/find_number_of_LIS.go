package one

func FindNumberOfLIS(nums []int) int {
	// 最大长度
	ls := make([]int, len(nums))
	// 最大长度对应的数量
	ns := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		mi := -1
		n := 0
		for j := i - 1; j >= 0; j-- {
			if nums[i] <= nums[j] {
				continue
			}
			if mi == -1 || ls[j] > ls[mi] {
				mi = j
				n = ns[j]
				continue
			}
			// 长度一致，那么数量需要叠加
			if ls[j] == ls[mi] {
				n += ns[j]
			}
		}
		if mi == -1 {
			ls[i] = 1
			ns[i] = 1
			continue
		}
		ls[i] = ls[mi] + 1
		ns[i] = n
	}
	ml := 0
	for _, l := range ls {
		if l > ml {
			ml = l
		}
	}
	ans := 0
	for i, l := range ls {
		if l == ml {
			ans += ns[i]
		}
	}
	return ans
}
