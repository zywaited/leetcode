package one

func Convert(s string, numRows int) string {
	l := len(s)
	if l <= numRows || numRows < 2 {
		return s
	}

	// 每组字符个数
	per := 2 * (numRows - 1)
	var zs []byte
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < l; j += per {
			zs = append(zs, s[j+i])
			// 两列中间元素
			// 第一行和最后一行中间是没有数据的，并且判断长度
			if i == 0 || i == numRows-1 || j+per-i >= l {
				continue
			}

			zs = append(zs, s[j+per-i])
		}
	}

	return string(zs)
}
