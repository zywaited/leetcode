package one

func SmallestSubsequence(s string, k int, letter byte, repetition int) string {
	cn := 0
	for i := range s {
		if s[i] == letter {
			cn++
		}
	}
	bs := make([]byte, 0, k)
	lc := 0
	for i := range s {
		// 当前字符优先级比前一个小，保证剩余字符不小于k-len(bs)，然后分情况
		// 1: 如果bs最后一个字符不是letter，那么直接删除就行了
		// 2: 如果bs最后一个字符是letter，那么判断lc+cn不小于repetition就可以删除
		for len(bs) > 0 && s[i] < bs[len(bs)-1] && k <= len(s)-i+len(bs)-1 && (bs[len(bs)-1] != letter || repetition <= lc-1+cn) {
			if bs[len(bs)-1] == letter {
				lc--
			}
			bs = bs[:len(bs)-1]
		}
		// bs的长度为k，但此时lc+cn == repetition，那么代表bs里面可能删除cn个字符来保证整体的letter长度满足
		if len(bs) == k && lc+cn == repetition {
			rn := 0 // 已经删除的非letter的数量
			rc := 0 // 已经删除的letter的数量
			for ; len(bs) > 0 && rn < cn; bs = bs[:len(bs)-1] {
				if bs[len(bs)-1] != letter {
					rn++
					continue
				}
				rc++
			}
			// letter被删除要补回去
			for ; rc > 0; rc-- {
				bs = append(bs, letter)
			}
		}
		if s[i] == letter {
			cn--
		}
		if len(bs) == k {
			continue
		}
		if s[i] == letter {
			lc++
		}
		bs = append(bs, s[i])
	}
	return string(bs)
}
