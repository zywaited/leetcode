package two

import "sort"

func MaxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] == envelopes[j][0] {
			// 高度大的放前面，这样从二维变成一维
			// 比如: [1,3][2,2][2,1]
			// 这样就不用考虑宽度了【高度比前一个大，那么宽度也一定比前一个大】
			return envelopes[i][1] > envelopes[j][1]
		}
		return envelopes[i][0] < envelopes[j][0]
	})
	indexes := make([]int, 0, len(envelopes))
	for _, envelope := range envelopes {
		if len(indexes) == 0 || envelope[1] > indexes[len(indexes)-1] {
			indexes = append(indexes, envelope[1])
			continue
		}
		// 二分替换
		s, e := 0, len(indexes)-1
		i := 0
		for s <= e {
			m := s + (e-s)>>1
			i = m
			if indexes[m] == envelope[1] {
				break
			}
			if indexes[m] < envelope[1] {
				s++
				continue
			}
			e--
		}
		indexes[i] = envelope[1]
	}
	return len(indexes)
}
