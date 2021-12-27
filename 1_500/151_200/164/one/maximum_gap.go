package one

func MaximumGap(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	// 桶排序
	min, max := nums[0], nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	// 按照个数等差计算，那么最大间距一定大于等于等差值
	// 那么每个桶的最大差值一定小于等差，即最大差值一定在桶与桶之间，每个桶只需保留最大最小值
	// 如果只有一个桶那么代表等差为0
	diff := (max - min) / (len(nums) - 1)
	if diff == 0 {
		diff = 1
	}
	size := (max-min)/diff + 1     // 这里+1是因为最大差值刚好就是整体
	buckets := make([][]int, size) // 0: 最小值, 1: 最大值
	for _, num := range nums {
		index := (num - min) / diff
		if len(buckets[index]) == 0 {
			buckets[index] = []int{num, num}
			continue
		}
		if num < buckets[index][0] {
			buckets[index][0] = num
		}
		if num > buckets[index][1] {
			buckets[index][1] = num
		}
	}
	ans := 0
	prev := buckets[0][1]
	for index := 1; index < len(buckets); index++ {
		if len(buckets[index]) > 0 {
			if buckets[index][0]-prev > ans {
				ans = buckets[index][0] - prev
			}
			prev = buckets[index][1]
		}
	}
	return ans
}
