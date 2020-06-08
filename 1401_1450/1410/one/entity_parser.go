package one

var relations = map[string]byte{
	"&quot;":  '"',
	"&apos;":  '\'',
	"&amp;":   '&',
	"&gt;":    '>',
	"&lt;":    '<',
	"&frasl;": '/',
}

func EntityParser(text string) string {
	ans := make([]byte, 0, len(text))
	// &的位置
	si := -1
	for index := range text {
		ans = append(ans, text[index])
		if si == -1 && text[index] != '&' {
			continue
		}
		// 重置&位置
		if text[index] == '&' {
			si = index
			continue
		}
		// 直到找到分号
		if text[index] != ';' {
			continue
		}
		// 不是特殊转换则重置&索引
		if relations[text[si:index+1]] == 0 {
			si = -1
			continue
		}
		// 先把加入的还原, 然后加入原始数据
		ans = ans[:len(ans)-(index-si)-1]
		ans = append(ans, relations[text[si:index+1]])
	}
	return string(ans)
}
