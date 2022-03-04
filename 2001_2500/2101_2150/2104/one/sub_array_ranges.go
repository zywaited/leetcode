package one

func SubArrayRanges(nums []int) int64 {
	ans := int64(0)
	var min []int
	var max []int
	curr := int64(0)
	for index, num := range nums {
		for len(max) > 0 && num <= nums[max[len(max)-1]] {
			numIndex := max[len(max)-1]
			if len(max) > 1 {
				numIndex -= max[len(max)-2] + 1
			}
			curr += int64(nums[max[len(max)-1]]-num) * int64(numIndex+1)
			max = max[:len(max)-1]
		}
		for len(min) > 0 && nums[min[len(min)-1]] <= num {
			numIndex := min[len(min)-1]
			if len(min) > 1 {
				numIndex -= min[len(min)-2] + 1
			}
			curr += int64(num-nums[min[len(min)-1]]) * int64(numIndex+1)
			min = min[:len(min)-1]
		}
		max = append(max, index)
		min = append(min, index)
		ans += curr
	}
	return ans
}
