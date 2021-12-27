package two

import (
	"fmt"
	"sort"
)

type SummaryRanges struct {
	data [][]int
}

func Constructor() SummaryRanges {
	return SummaryRanges{}
}

// copy 可改为 有序集合【跳表或红黑树】
func (sr *SummaryRanges) AddNum(val int) {
	fmt.Println(val)
	index := sort.Search(len(sr.data), func(i int) bool {
		return sr.data[i][1] >= val
	})
	if index == len(sr.data) {
		if index > 0 && sr.data[index-1][1]+1 == val {
			sr.data[index-1][1] = val
			return
		}
		sr.data = append(sr.data, []int{val, val})
		return
	}
	if sr.data[index][0] <= val || sr.data[index][1] == val {
		return
	}
	if sr.data[index][0]-1 == val {
		sr.data[index][0] = val
		if index >= 1 && sr.data[index-1][1]+1 == val {
			sr.data[index-1][1] = sr.data[index][1]
			copy(sr.data[index:], sr.data[index+1:])
			sr.data = sr.data[:len(sr.data)-1]
		}
		return
	}
	if index >= 1 && sr.data[index-1][1]+1 == val {
		sr.data[index-1][1] = val
		return
	}
	data := make([][]int, len(sr.data)+1)
	copy(data[:], sr.data[:index])
	data[index] = []int{val, val}
	copy(data[index+1:], sr.data[index:])
	sr.data = data
}

func (sr *SummaryRanges) GetIntervals() [][]int {
	return sr.data
}
