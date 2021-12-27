package one

// 做该题的时候翻译是有问题，相同顺序也是可以的，具体可看给的例子4或者英文版
func MinSteps(s string, t string) int {
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}
	num := 0
	for i := range t {
		if m[t[i]] > 0 {
			m[t[i]]--
			continue
		}
		num++
	}
	return num
}
