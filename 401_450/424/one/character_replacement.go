package one

func CharacterReplacement(s string, k int) int {
	cnt := [26]int{}
	left := 0
	mn := 0
	for right, c := range s {
		cnt[c-'A']++
		mn = max(mn, cnt[c-'A'])
		if right-left+1-mn > k {
			cnt[s[left]-'A']--
			left++
		}
	}
	return len(s) - left
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
