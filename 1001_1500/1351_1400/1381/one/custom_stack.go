package one

type CustomStack struct {
	// 模拟栈
	stack []int
	max   int
	// 为了操作O(1)
	sum map[int]int
}

func Constructor(maxSize int) CustomStack {
	return CustomStack{
		stack: make([]int, 0, maxSize),
		max:   maxSize,
		sum:   make(map[int]int),
	}
}

func (cs *CustomStack) Push(x int) {
	if len(cs.stack) == cs.max {
		return
	}
	cs.stack = append(cs.stack, x)
}

func (cs *CustomStack) Pop() int {
	if len(cs.stack) == 0 {
		return -1
	}
	l := len(cs.stack)
	num := cs.stack[l-1]
	cs.stack = cs.stack[:l-1]
	if cs.sum[l] == 0 {
		return num
	}
	// 获取并清除当前可加数据
	num += cs.sum[l]
	cs.sum[l-1] += cs.sum[l]
	cs.sum[l] = 0
	return num
}

func (cs *CustomStack) Increment(k int, val int) {
	// 长度规划
	if k > len(cs.stack) {
		k = len(cs.stack)
	}
	cs.sum[k] += val
}
