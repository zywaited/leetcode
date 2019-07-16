package two

func MajorityElement(nums []int) int {
	mp := make(map[int]int)
	half := len(nums) >> 1
	for _, num := range nums {
		mp[num]++
		if mp[num] <= half {
			continue
		}
		return num
	}
	// 众数总是存在，不会走到这里
	return -1
}
