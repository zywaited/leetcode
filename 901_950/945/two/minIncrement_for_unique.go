package two

func MinIncrementForUnique(A []int) int {
	// 统计每个数量
	nums := make(map[int]int, len(A)<<1)
	max := 0
	for _, num := range A {
		nums[num]++
		if num > max {
			max = num
		}
	}
	ans := 0
	// 这个为了重复利用，按顺序查找
	for num := 0; num <= max; num++ {
		if nums[num] > 1 {
			// 把多余数量统统移到下一位
			ans += nums[num] - 1
			nums[num+1] += nums[num] - 1
		}
	}
	// 最后把max+1的数据处理一下
	if nums[max+1] > 1 {
		// 1+2+3+4.....
		ans += ((nums[max+1] - 1) * nums[max+1]) >> 1
	}
	return ans
}
