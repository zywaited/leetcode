package one

func MissingNumber(nums []int) int {
	// 进行预处理
	// 以值为索引
	nl := len(nums)
	max := false
	num := -1
	for _, num = range nums {
		if num < 0 {
			num = -num
		}
		if num >= nl {
			max = true
			continue
		}
		if nums[num] < 0 {
			continue
		}
		// 0需要特殊处理，如果存在则把值置为一个下标不存在数字即可
		if nums[num] == 0 {
			if nums[0] > 0 {
				nums[0] = -nums[0]
			}
			nums[num] = -nl
			continue
		}
		nums[num] = -nums[num]
	}
	if !max {
		return nl
	}
	r := 0
	for _, num = range nums {
		if num >= 0 {
			break
		}
		r++
	}
	return r
}
