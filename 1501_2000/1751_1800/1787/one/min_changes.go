package one

func MinChanges(nums []int, k int) int {
	// 每层对应的数字及数量
	kn := make([]map[int]int, k)
	// 每层的数量
	ktn := make([]int, k)
	for i := 0; i < len(nums); i++ {
		if i < k {
			kn[i] = make(map[int]int)
		}
		kn[i%k][nums[i]]++
		ktn[i%k]++
	}
	curr := make([]int, 1<<10)
	// 当前层变换为某个数字所需的最小次数
	pm := len(nums)
	for i := range curr {
		curr[i] = ktn[0] - kn[0][i]
		pm = min(pm, curr[i])
	}
	next := make([]int, 1<<10)
	npm := 0
	for i := 1; i < k; i++ {
		npm = len(nums)
		for nn := 0; nn < 1<<10; nn++ {
			next[nn] = pm + ktn[i]  // 默认改变当前层的数字
			for pn := range kn[i] { // 不改变当前层的数字
				next[nn] = min(next[nn], curr[nn^pn]+ktn[i]-kn[i][pn])
			}
			npm = min(npm, next[nn])
		}
		curr, next = next, curr
		pm = npm
	}
	return curr[0]
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
