package two

func MinMovesToMakePalindrome(s string) int {
	// 在对一种方法中，计算了每种字符的移动次数，选择了最小的一个
	// 这里就不用计算最小的，因为不管怎么移动总次数是不变的
	cis := [26][]int{}
	for i := range s {
		cis[s[i]-'a'] = append(cis[s[i]-'a'], i+1)
	}
	tr := make(tree, len(s)+1)
	ans := 0
	ci := 1
	si := 0
	for times := len(s) / 2; times > 0; times-- {
		for ; si < len(s) && len(cis[s[si]-'a']) <= 1; si++ {
		}
		i := s[si] - 'a'
		li := cis[i][0] + tr.sum(cis[i][0])
		ri := cis[i][len(cis[i])-1] + tr.sum(cis[i][len(cis[i])-1])
		ans += li - ci + (len(s) - ci + 1) - (ri)
		if ri < len(s)-ci+1 {
			tr.update(cis[i][len(cis[i])-1], -1)
		}
		cis[i] = cis[i][1 : len(cis[i])-1]
		ci++
		si++
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
