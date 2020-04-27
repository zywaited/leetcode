package one

func GetTriggerTime(increase [][]int, requirements [][]int) []int {
	ps := make([][3]int, 0, len(increase)+1)
	ps = append(ps, [3]int{})
	// 总和
	for index, incr := range increase {
		ps = append(ps, [3]int{ps[index][0] + incr[0], ps[index][1] + incr[1], ps[index][2] + incr[2]})
	}
	ans := make([]int, len(requirements))
	for index, requirement := range requirements {
		// 默认不能完成
		ans[index] = -1
		s := 0
		e := len(increase)
		for s <= e {
			m := s + (e-s)>>1
			if ps[m][0] >= requirement[0] && ps[m][1] >= requirement[1] && ps[m][2] >= requirement[2] {
				ans[index] = m
				e = m - 1
				continue
			}
			s = m + 1
		}
	}
	return ans
}
