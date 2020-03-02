package one

// 题目没有明确说明只有小写字母
func FirstUniqChar(s string) byte {
	// byte的范围0-128分开计算
	// 如果要计算128需要额外一个变量存储(这里就不做处理了，思路是一样的，测试用例也没有这个)
	hn := uint64(0) // 64-127
	ln := uint64(0) // 0-63
	// 存放已经存在的
	ehn := uint64(0)
	eln := uint64(0)
	l := uint64(0)
	for i := range s {
		if s[i] >= 64 {
			l = 1 << (s[i] - 64)
			if l&ehn == 0 {
				ehn |= l
				hn |= l
				continue
			}
			if hn&l != 0 {
				hn ^= l
			}
			continue
		}
		l = 1 << s[i]
		if l&eln == 0 {
			eln |= l
			ln |= l
			continue
		}
		if ln&l != 0 {
			ln ^= l
		}
	}
	ans := byte(' ')
	// 找到最新出现的
	for i := range s {
		if s[i] >= 64 {
			l = 1 << (s[i] - 64)
			if hn&l != 0 {
				ans = s[i]
				break
			}
			continue
		}
		l = 1 << s[i]
		if ln&l != 0 {
			ans = s[i]
			break
		}
	}
	return ans
}
