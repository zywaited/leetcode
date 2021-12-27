package two

func Combine(n int, k int) [][]int {
	nums := make([]int, 0, k)
	sn := 1
	for ; sn <= k; sn++ {
		nums = append(nums, sn)
	}
	// 对nums的前K个数作变动：例如例子n=4,k=2
	// 取到[1,4]时需要对把前一位加1，后续都在自身前一位值加1，也就是变成了[2,3]继续操作直到[2,4]
	rs := make([][]int, 0)
	index := k - 1 // 当前位索引
	for index >= 0 {
		rs = append(rs, append(make([]int, 0, k), nums...))
		if nums[index] < n {
			// 最后一位增加
			nums[index]++
			continue
		}
		// 已经到底，那就前一位
		// 判断是否可变动前一位并且满足k个数
		for ; index > 0; index-- {
			nums[index-1]++ // 前一位加1
			if nums[index-1] >= n || (index+n-nums[index-1]) < k {
				continue
			}
			break
		}
		// 索引不能为0
		if index == 0 {
			break
		}
		// 后面位为前一位加1
		for ; index < k; index++ {
			nums[index] = nums[index-1] + 1
		}
		index = k - 1
	}
	return rs
}
