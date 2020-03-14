package three

// 贪心+二分
// 尽量每个数字小
func LengthOfLIS(nums []int) int {
	ls := make([]int, 0, len(nums))
	bs := func(target int) int {
		s, e := 0, len(ls)-1
		index := 0
		for s <= e {
			m := s + (e-s)>>1
			if ls[m] == target {
				index = m
				break
			}
			if ls[m] < target {
				s = m + 1
				continue
			}
			e = m - 1
			index = m
		}
		return index
	}
	for _, num := range nums {
		if len(ls) == 0 || ls[len(ls)-1] < num {
			ls = append(ls, num)
		}
		ls[bs(num)] = num
	}
	return len(ls)
}
