package one

func takeCharacters(s string, k int) int {
	cnt := make([]int, 3)
	ans := len(s)

	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}

	if cnt[0] >= k && cnt[1] >= k && cnt[2] >= k {
		ans = min(ans, len(s))
	} else {
		return -1
	}

	l := 0
	for r := 0; r < len(s); r++ {
		cnt[s[r]-'a']--
		for l < r && (cnt[0] < k || cnt[1] < k || cnt[2] < k) {
			cnt[s[l]-'a']++
			l++
		}
		if cnt[0] >= k && cnt[1] >= k && cnt[2] >= k {
			ans = min(ans, len(s)-(r-l+1))
		}
	}

	return ans
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
