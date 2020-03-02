package two

// 题目没有明确说明只有小写字母
func FirstUniqChar(s string) byte {
	nums := make([]int, 129) // 0-128
	for i := range s {
		nums[s[i]]++
	}
	for i := range s {
		if nums[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}
