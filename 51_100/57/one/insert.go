package one

func Insert(intervals [][]int, newInterval []int) [][]int {
	rs := make([][]int, 0)
	if len(intervals) == 0 {
		rs = append(rs, newInterval)
		return rs
	}
	insert := false
	for i, interval := range intervals {
		// 判断是否需要合并区间
		if interval[1] < newInterval[0] {
			rs = append(rs, interval)
			continue
		}
		if interval[0] > newInterval[1] {
			rs = append(rs, newInterval)
			rs = append(rs, intervals[i:]...)
			insert = true
			break
		}
		if interval[0] < newInterval[0] {
			newInterval[0] = interval[0]
		}
		if interval[1] > newInterval[1] {
			newInterval[1] = interval[1]
		}
	}
	if !insert {
		rs = append(rs, newInterval)
	}
	return rs
}
