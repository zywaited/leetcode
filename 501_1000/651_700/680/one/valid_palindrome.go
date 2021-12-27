package one

func ValidPalindrome(s string) bool {
	l, r := 0, len(s)-1
	removed := false   // 删除是否都已经试过了
	rli, rri := -1, -1 // 移除时候的左右索引
	for l < r {
		if s[l] == s[r] {
			l++
			r--
			continue
		}
		if removed {
			return false
		}
		if rli > -1 {
			// 恢复
			l = rli
			r = rri
		}
		if rli == -1 && s[l+1] == s[r] {
			rli = l
			rri = r
			l += 2
			r--
			continue
		}
		if s[l] == s[r-1] {
			l++
			r -= 2
			removed = true
			continue
		}
		return false
	}
	return true
}
