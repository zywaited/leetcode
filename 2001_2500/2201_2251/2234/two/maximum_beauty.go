package two

import "sort"

func MaximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	sort.Ints(flowers)
	ti := sort.SearchInts(flowers, target)
	sums := make([]int64, ti+1)
	for i := 0; i < ti; i++ {
		sums[i+1] = sums[i] + int64(flowers[i])
	}
	ans := int64(0)
	en := newFlowers
	for i := ti; i >= 0 && (i == len(flowers) || int64(target-flowers[i]) <= en); i-- {
		if i == 0 {
			ans = max(ans, int64(full)*int64(len(flowers)))
			continue
		}
		if i < len(flowers) && flowers[i] < target {
			en -= int64(target - flowers[i])
		}
		s := 0
		e := target - 1
		for s <= e {
			m := s + (e-s)>>1
			mi := sort.SearchInts(flowers[:i], m)
			if int64(m)*int64(mi)-sums[mi] <= en {
				s = m + 1
				continue
			}
			e = m - 1
		}
		ans = max(ans, int64(full)*int64(len(flowers)-i)+int64(partial)*int64(e))
	}
	return ans
}

func max(f, s int64) int64 {
	if f >= s {
		return f
	}
	return s
}
