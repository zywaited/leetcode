package two

// 做该题的时候翻译是有问题，相同顺序也是可以的，具体可看给的例子4或者英文版
func MinSteps(s string, t string) int {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
		m[t[i]]--
	}
	num := 0
	for _, n := range m {
		// > 0 代表s中有剩余的需要替换掉t
		if n > 0 {
			num += n
		}
	}
	return num
}
