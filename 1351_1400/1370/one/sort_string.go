package one

func SortString(s string) string {
	bs := make([]int, 26)
	for i := range s {
		bs[s[i]-'a']++
	}
	ans := make([]byte, 0, len(s))
	for len(ans) < len(s) {
		// 升序
		for i := 0; i < 26; i++ {
			if bs[i] > 0 {
				ans = append(ans, byte(i)+'a')
				bs[i]--
			}
		}
		// 降序
		for i := 25; i >= 0; i-- {
			if bs[i] > 0 {
				ans = append(ans, byte(i)+'a')
				bs[i]--
			}
		}
	}
	return string(ans)
}
