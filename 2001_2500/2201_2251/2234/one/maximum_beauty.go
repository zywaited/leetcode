package one

import "sort"

func MaximumBeauty(flowers []int, newFlowers int64, target int, full int, partial int) int64 {
	sort.Ints(flowers)
	ti := sort.SearchInts(flowers, target)
	tn := len(flowers) - ti
	flowers = flowers[:ti]
	sums := make([]int64, len(flowers)+1)
	for i := range flowers {
		sums[i+1] = sums[i] + int64(flowers[i])
	}
	ans := int64(0)
	for t := 1; t <= target; t++ {
		i := sort.SearchInts(flowers, t+1)
		n := newFlowers
		if sums[i] < int64(t)*int64(i) {
			n -= int64(t)*int64(i) - sums[i]
			if n < 0 {
				break
			}
		}
		if t == target {
			ans = max(ans, int64(full)*int64(len(flowers)+tn))
			continue
		}
		if i == 0 {
			continue
		}
		j := sort.Search(len(flowers), func(j int) bool {
			s := int64(0)
			if j < i {
				s = sums[len(flowers)] - sums[i] + int64(t)*int64(i-j)
			} else {
				s = sums[len(flowers)] - sums[j]
			}
			return j > 0 && int64(target)*int64(len(flowers)-j)-s <= n
		})
		ans = max(ans, int64(full)*int64(len(flowers)-j+tn)+int64(partial)*int64(t))
	}
	return ans
}

func max(f, s int64) int64 {
	if f >= s {
		return f
	}
	return s
}
