package one

import "container/heap"

func ReorganizeString(S string) string {
	max := 0
	ns := make(nh, 26)
	for index := range ns {
		ns[index] = &node{
			char: byte(index) + 'a',
			num:  0,
		}
	}
	for index := range S {
		ns[S[index]-'a'].num++
		if ns[S[index]-'a'].num > max {
			max = ns[S[index]-'a'].num
		}
	}
	if max > (len(S)+1)>>1 {
		return ""
	}
	heap.Init(&ns)
	bs := make([]byte, 0, len(S))
	for len(bs) < len(S) {
		// 每次取出前两种字符
		f, s := heap.Pop(&ns).(*node), heap.Pop(&ns).(*node)
		if f.num > 0 {
			bs = append(bs, f.char)
			f.num--
		}
		if s.num > 0 {
			bs = append(bs, s.char)
			s.num--
		}
		heap.Push(&ns, f)
		heap.Push(&ns, s)
	}
	return string(bs)
}

type (
	node struct {
		char byte
		num  int
	}
	nh []*node
)

func (h *nh) Len() int {
	return len(*h)
}

func (h *nh) Less(i, j int) bool {
	return (*h)[i].num > (*h)[j].num
}

func (h *nh) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *nh) Push(n interface{}) {
	*h = append(*h, n.(*node))
}

func (h *nh) Pop() interface{} {
	n := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return n
}
