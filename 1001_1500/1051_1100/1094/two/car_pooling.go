package two

import "sort"

func CarPooling(trips [][]int, capacity int) bool {
	sort.Sort(Trips(trips))               // 先按照上车点排序
	indexes := make([]int, 0, len(trips)) // 已经上车的索引
	minIndex := -1                        // 下车点最小值
	search := func(s int) []int {         // 找到下车的索引
		if s >= minIndex && minIndex >= 0 {
			l, r := 0, len(indexes)-1
			for l < r {
				for l < r && trips[indexes[l]][2] <= s {
					l++
				}
				for r >= l && trips[indexes[r]][2] > s {
					if minIndex <= s || trips[indexes[r]][2] < minIndex {
						minIndex = trips[indexes[r]][2]
					}
					r--
				}
				if l < r {
					indexes[l], indexes[r] = indexes[r], indexes[l]
					if minIndex <= s || trips[indexes[r]][2] < minIndex {
						minIndex = trips[indexes[r]][2]
					}
					l++
					r--
				}
			}
			i := l
			if trips[indexes[l]][2] <= s {
				i++
			}
			if i == len(indexes) {
				minIndex = -1
			}
			dn := indexes[:i]
			indexes = indexes[i:]
			return dn
		}
		return nil
	}
	for i, trip := range trips {
		for _, d := range search(trip[1]) { // 恢复座位数
			capacity += trips[d][0]
		}
		if capacity < trip[0] {
			return false
		}
		capacity -= trip[0]
		indexes = append(indexes, i)
		if minIndex < 0 || trip[2] < minIndex {
			minIndex = trip[2]
		}
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
