package two

func DailyTemperatures(T []int) []int {
	ans := make([]int, len(T))
	// 用于保存索引
	indexes := make([]int, 0, len(T))
	indexes = append(indexes, len(T)-1)
	// 通过二分法查找替换元素
	bs := func(i int) int {
		s, e := 0, len(indexes)-1
		ri := len(indexes) // 默认不替换
		for s <= e {
			m := s + (e-s)>>1
			if T[indexes[m]] <= T[i] {
				ri = m
				e = m - 1
				continue
			}
			s = m + 1
		}
		return ri
	}
	for i := len(T) - 2; i >= 0; i-- {
		indexes = indexes[:bs(i)]
		if len(indexes) > 0 {
			ans[i] = indexes[len(indexes)-1] - i
		}
		indexes = append(indexes, i)
	}
	return ans
}
