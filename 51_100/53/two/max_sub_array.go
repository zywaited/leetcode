package two

func MaxSubArray(nums []int) int {
	// 贪心
	// 如果区间和是正数，就贪更多可能的数相加
	// 但是需要保留当前区间的最大值
	max := nums[0]
	sum := 0
	for _, num := range nums {
		if sum > 0 {
			sum += num
		} else {
			sum = num // 小于等于0区间的起点发生变化
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
