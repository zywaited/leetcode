package one

func FullJustify(words []string, maxWidth int) []string {
	ans := make([]string, 0)
	cl := 0  // 当前行的单词长度
	cwl := 0 // 当前的单词个数
	uel := 0 // 截止上一行使用的单词个数
	cb := make([]byte, 0)
	tmp := 0
	for index, word := range words {
		wl := len(word)
		if cl+wl+cwl > maxWidth {
			// 当前的不能用
			// 只有一个单词后面全是空格
			if cwl == 1 {
				for tmp = maxWidth - cl - cwl; tmp > 0; tmp-- {
					cb = append(cb, ' ')
				}
			} else {
				// 判断空格是否均分
				bs, es := (maxWidth-cl)/(cwl-1), (maxWidth-cl)%(cwl-1)
				cb = cb[:0]
				for tmp = uel; tmp <= index-2; tmp++ {
					cb = append(cb, words[tmp]...)
					if es > 0 {
						cb = append(cb, ' ')
						es--
					}
					for tbs := bs; tbs > 0; tbs-- {
						cb = append(cb, ' ')
					}
				}
				cb = append(cb, words[index-1]...) // 写入最后的那个数据
			}
			uel += cwl
			cl = 0
			cwl = 0
			ans = append(ans, string(cb))
			cb = cb[:0]
		}
		cb = append(cb, word...)
		if cl+wl+cwl == maxWidth {
			ans = append(ans, string(cb))
			cl = 0
			uel += cwl + 1
			cwl = 0
			cb = cb[:0]
			continue
		}
		// 是否是最后一行
		if index == len(words)-1 {
			for tmp = maxWidth - wl - cl - cwl; tmp > 0; tmp-- {
				cb = append(cb, ' ')
			}
			ans = append(ans, string(cb))
			break // 结束了
		}
		cb = append(cb, ' ')
		cl += wl
		cwl++
	}
	return ans
}
