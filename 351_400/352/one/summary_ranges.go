package one

import "sort"

type SummaryRanges struct {
	data []int
	sort bool
}

func Constructor() SummaryRanges {
	return SummaryRanges{}
}

func (sr *SummaryRanges) AddNum(val int) {
	sr.data = append(sr.data, val)
	sr.sort = false
}

func (sr *SummaryRanges) GetIntervals() [][]int {
	if !sr.sort {
		sort.Ints(sr.data)
		sr.sort = true
	}
	ans := make([][]int, 0, len(sr.data))
	for _, num := range sr.data {
		if len(ans) == 0 || ans[len(ans)-1][1]+1 < num {
			ans = append(ans, []int{num, num})
			continue
		}
		ans[len(ans)-1][1] = num
	}
	return ans
}
