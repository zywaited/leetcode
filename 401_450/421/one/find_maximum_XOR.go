package one

type node [2]*node

func FindMaximumXOR(nums []int) int {
	root := &node{}
	insert := func(num int) {
		c := root
		for i := 30; i >= 0; i-- {
			b := num >> uint(i) & 1
			if c[b] == nil {
				c[b] = &node{}
			}
			c = c[b]
		}
	}
	for _, num := range nums {
		insert(num)
	}
	find := func(num int) int {
		c := root
		v := 0
		for i := 30; i >= 0; i-- {
			b := num >> uint(i) & 1
			if c[b^1] != nil {
				c = c[b^1]
				v += 1 << uint(i)
				continue
			}
			c = c[b]
		}
		return v
	}
	ans := 0
	for _, num := range nums {
		ans = max(ans, find(num))
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
