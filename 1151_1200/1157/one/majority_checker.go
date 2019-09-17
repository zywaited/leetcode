package one

// 最容易想到就是把每个数字归类
// 然后判断数量是否满足
type MajorityChecker struct {
	numsIndexes map[int][]int
}

func Constructor(arr []int) MajorityChecker {
	mc := MajorityChecker{numsIndexes: make(map[int][]int)}
	for index, num := range arr {
		mc.numsIndexes[num] = append(mc.numsIndexes[num], index)
	}
	return mc
}

func (mc *MajorityChecker) Query(left int, right int, threshold int) int {
	// 二分查找边界
	for num, indexes := range mc.numsIndexes {
		if len(indexes) < threshold || indexes[0] > right || indexes[len(indexes)-1] < left {
			continue
		}
		l := mc.searchIndex(num, 0, len(indexes)-1, left, false)
		r := mc.searchIndex(num, l, len(indexes)-1, right, true)
		if r-l+1 >= threshold {
			return num
		}
	}
	return -1
}

func (mc *MajorityChecker) searchIndex(num, l, r, t int, mr bool) int {
	indexes := mc.numsIndexes[num]
	if indexes[l] == t {
		return l
	}
	if indexes[r] == t {
		return r
	}
	m := 0
	rs := 0
	for l <= r {
		m = l + (r-l)>>1
		if indexes[m] == t {
			return m
		}
		if indexes[m] < t {
			l = m + 1
			if mr {
				rs = m
			}
			continue
		}
		r = m - 1
		if !mr {
			rs = m
		}
	}
	return rs
}
