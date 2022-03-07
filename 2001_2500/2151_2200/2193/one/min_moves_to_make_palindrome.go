package one

func MinMovesToMakePalindrome(s string) int {
	cis := [26][]int{}
	for i := range s {
		cis[s[i]-'a'] = append(cis[s[i]-'a'], i+1)
	}
	tr := make(tree, len(s)+1)
	ans := 0
	ci := 1
	for times := len(s) / 2; times > 0; times-- {
		step := 0
		si := -1
		li := 0
		ri := 0
		for i := range cis {
			if len(cis[i]) <= 1 {
				continue
			}
			tli := cis[i][0] + tr.sum(cis[i][0])
			tri := cis[i][len(cis[i])-1] + tr.sum(cis[i][len(cis[i])-1])
			cs := tli - ci + (len(s) - ci + 1) - (tri)
			if si < 0 || cs < step {
				si = i
				step = cs
				li = tli
				ri = tri
			}
		}
		if ci < li {
			tr.update(1, 1)
			tr.update(cis[si][0], -1)
		}
		if ri < len(s)-ci+1 {
			tr.update(cis[si][len(cis[si])-1], -1)
		}
		cis[si] = cis[si][1 : len(cis[si])-1]
		ans += step
		ci++
	}
	return ans
}

type tree []int

func (t *tree) len() int {
	return len(*t)
}

func (t *tree) lowBit(index int) int {
	return index & (-index)
}

func (t *tree) update(index, value int) {
	for ; index < t.len(); index += t.lowBit(index) {
		(*t)[index] = (*t)[index] + value
	}
}

func (t *tree) sum(index int) int {
	sum := 0
	for ; index > 0; index -= t.lowBit(index) {
		sum += (*t)[index]
	}
	return sum
}
