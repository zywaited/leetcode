package one

import "sort"

// 力扣有一个测试用例结果是正确的但是无法通过，具体可看测试文件最后一个
func CarPooling(trips [][]int, capacity int) bool {
	sort.Sort(Trips(trips))             // 先按照上车点排序
	index := make([]int, 0, len(trips)) // 已经上车的索引
	dn := func(s int) []int {           // 找到下车的索引
		l, r := 0, len(index)-1
		i := 0
		for l <= r {
			m := l + (r-l)>>1
			if trips[index[m]][2] <= s {
				i = l + 1
				l = m + 1
				continue
			}
			r = m - 1
		}
		ds := index[:i]
		index = index[i:]
		return ds
	}
	for i, trip := range trips {
		for _, d := range dn(trip[1]) { // 恢复座位数
			capacity += trips[d][0]
		}
		if capacity < trip[0] {
			return false
		}
		capacity -= trip[0]
		index = append(index, i)
	}
	return true
}

type Trips [][]int

func (ts Trips) Len() int {
	return len(ts)
}

func (ts Trips) Less(i, j int) bool {
	return ts[i][1] <= ts[j][1]
}

func (ts Trips) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}
