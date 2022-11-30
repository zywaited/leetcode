package one

type FreqStack struct {
	vm map[int]int
	nm map[int][]int
	m  int
}

func Constructor() FreqStack {
	return FreqStack{vm: map[int]int{}, nm: map[int][]int{}}
}

func (fs *FreqStack) Push(val int) {
	fs.vm[val]++
	fs.nm[fs.vm[val]] = append(fs.nm[fs.vm[val]], val)
	if fs.m < fs.vm[val] {
		fs.m = fs.vm[val]
	}
}

func (fs *FreqStack) Pop() int {
	val := fs.nm[fs.m][len(fs.nm[fs.m])-1]
	fs.nm[fs.m] = fs.nm[fs.m][:len(fs.nm[fs.m])-1]
	if len(fs.nm[fs.m]) == 0 {
		fs.m--
	}
	fs.vm[val]--
	return val
}
