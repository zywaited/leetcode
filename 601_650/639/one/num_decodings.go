package one

const mod = 1e9 + 7

// 为了方便就用if else大法
func NumDecodings(s string) int {
	ans := 1
	last := 0
	pn := uint8('0') // 前一个字符
	for i := range s {
		if s[i] != '*' {
			if s[i] == '0' {
				if pn == '0' || pn > '2' {
					return 0
				}
				if pn != '*' {
					ans = last
				} else {
					ans = (last << 1) % mod
				}
			} else if pn == '0' || pn > '2' || (pn == '2' && s[i] > '6') {
				last = ans
			} else if pn != '*' || (pn == '*' && s[i] > '6') {
				ans, last = (ans+last)%mod, ans
			} else {
				ans, last = (ans+last<<1)%mod, ans
			}
		} else if pn == '0' || pn > '2' {
			// 只能1-9
			ans, last = (ans*9)%mod, ans
		} else if pn == '2' {
			// 结合只能1-6
			ans, last = (ans*9+last*6)%mod, ans
		} else if pn != '*' {
			ans, last = (ans*9+last*9)%mod, ans
		} else {
			ans, last = (ans*9+last*15)%mod, ans
		}
		pn = s[i]
	}
	return ans
}
