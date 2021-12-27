package two

import "bytes"

func ShortestPalindrome(s string) string {
	l := len(s)
	if l < 1 {
		return ""
	}
	if l == 1 {
		return s
	}
	// 预处理
	buffer := &bytes.Buffer{}
	bs := []byte{'#'}
	for _, c := range s {
		bs = append(bs, byte(c))
		bs = append(bs, '#')
	}
	bsl := l*2 + 1
	ps := make([]int, bsl)
	// 从中心点往左
	center := l
	left := 0
	count := true
	border := 0
	for i := l; i >= 0; i-- {
		if i > left {
			// 对称性
			ps[i] = ps[2*center-i]
			if i-ps[i] < left {
				ps[i] = left - i
			}

			if i-ps[i] < left {
				count = false
			}
		}

		// 不需要计算
		if !count {
			continue
		}

		// 判断是否可以扩大回文
		for i-ps[i]-1 >= 0 && i+ps[i]+1 < bsl && bs[i-ps[i]-1] == bs[i+ps[i]+1] {
			ps[i]++
		}

		if i-ps[i] == 0 {
			// 已经是最优
			border = i
			break
		}

		// 继续往左偏移中心点
		center = i
		left = i - ps[i]
	}

	buffer.Reset()
	for i := bsl - 2; i > border+ps[border]; i -= 2 {
		buffer.WriteByte(bs[i])
	}
	buffer.WriteString(s)
	return buffer.String()
}
