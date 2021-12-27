package one

const (
	mod = 1e9 + 7
	max = 1e5
)

func CreateSortedArray(instructions []int) int {
	ans := 0
	t := newTree(max)
	for _, num := range instructions {
		ans = (ans + min(t.sum(num-1), t.sum(max)-t.sum(num))) % mod
		t.update(num)
	}
	return ans
}

type tree []int

func newTree(num int) *tree {
	t := make(tree, num+1)
	return &t
}

func (t *tree) len() int {
	return len(*t)
}

func (t *tree) lowBit(index int) int {
	return index & (-index)
}

func (t *tree) update(index int) {
	for ; index < t.len(); index += t.lowBit(index) {
		(*t)[index] = ((*t)[index] + 1) % mod
	}
}

func (t *tree) sum(index int) int {
	sum := 0
	for ; index > 0; index -= t.lowBit(index) {
		sum = (sum + (*t)[index]) % mod
	}
	return sum
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
