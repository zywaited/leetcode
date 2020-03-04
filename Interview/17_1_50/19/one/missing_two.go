package one

func MissingTwo(nums []int) []int {
	ans := []int{len(nums) + 1, len(nums) + 2} // 默认最后两位没有
	diff := 0                                  // 缺失数字数量
	// 以索引为依据判断数据
	for _, num := range nums {
		if num < 0 {
			num = -num
		}
		// 超范围
		if num > len(nums) {
			ans[num-len(nums)-1] = 0
			diff++
			continue
		}
		nums[num-1] = -nums[num-1]
	}
	if diff == 0 {
		return ans
	}
	// 找到nums中大于0的就是缺失的数字
	for index, num := range nums {
		if num < 0 {
			continue
		}
		if diff == 1 {
			if ans[0] == 0 {
				ans[0] = index + 1
			} else {
				ans[1] = index + 1
			}
			break
		}
		ans[0] = index + 1
		diff--
	}
	return ans
}
