package one

// This is the MountainArray's API interface.
// You should not implement it, or speculate about its implementation
type MountainArray struct {
}

func (ma *MountainArray) get(index int) int {
	return -1
}

func (ma *MountainArray) length() int {
	return -1
}

func FindInMountainArray(target int, mountainArr *MountainArray) int {
	s, e := 0, mountainArr.length()-1
	// 先找中间点
	for s < e {
		m := s + (e-s)>>1
		if mountainArr.get(m) < mountainArr.get(m+1) {
			s = m + 1
			continue
		}
		// 这里不减1，可能刚好是中间那个数据(大于所有数)
		e = m
	}
	// incr代表是否是递增
	bs := func(s, e int, incr bool) int {
		for s <= e {
			m := s + (e-s)>>1
			mv := mountainArr.get(m)
			if mv == target {
				return m
			}
			if mv < target {
				if incr {
					s = m + 1
				} else {
					e = m - 1
				}
				continue
			}
			if incr {
				e = m - 1
			} else {
				s = m + 1
			}
		}
		return -1
	}
	// 先查左边是否有数据
	index := bs(0, e, true)
	if index != -1 {
		return index
	}
	return bs(e+1, mountainArr.length()-1, false)
}
