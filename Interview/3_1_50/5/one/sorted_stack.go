package one

// Golang没有栈结构
// 用数组或者链表模拟栈操作
type SortedStack struct {
	stack []int
	help  []int
}

func Constructor() SortedStack {
	return SortedStack{}
}

func (this *SortedStack) Push(val int) {
	for len(this.stack) > 0 {
		if this.stack[len(this.stack)-1] >= val {
			break
		}
		this.help = append(this.help, this.stack[len(this.stack)-1])
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, val)
	for len(this.help) > 0 {
		this.stack = append(this.stack, this.help[len(this.help)-1])
		this.help = this.help[:len(this.help)-1]
	}
}

func (this *SortedStack) Pop() {
	if len(this.stack) > 0 {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *SortedStack) Peek() int {
	if len(this.stack) == 0 {
		return -1
	}
	return this.stack[len(this.stack)-1]
}

func (this *SortedStack) IsEmpty() bool {
	return len(this.stack) == 0
}
