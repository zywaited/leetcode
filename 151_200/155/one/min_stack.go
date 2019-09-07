package one

// 数组
type MinStack struct {
	data []int
	n    int
	min  []int
	mn   int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}

func (ms *MinStack) Push(x int) {
	ms.data = append(ms.data, x)
	if len(ms.min) == 0 || x <= ms.min[ms.mn-1] {
		ms.min = append(ms.min, x)
		ms.mn++
	}
	ms.n++
}

func (ms *MinStack) Pop() {
	if ms.n > 0 {
		if ms.data[ms.n-1] == ms.min[ms.mn-1] {
			ms.min = ms.min[:ms.mn-1]
			ms.mn--
		}
		ms.data = ms.data[:ms.n-1]
		ms.n--
	}

}

func (ms *MinStack) Top() int {
	if ms.n > 0 {
		return ms.data[ms.n-1]
	}
	return 0
}

func (ms *MinStack) GetMin() int {
	if ms.n > 0 {
		return ms.min[len(ms.min)-1]
	}
	return 0
}
