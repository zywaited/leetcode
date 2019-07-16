package two

import "bytes"

func LongestPalindrome(s string) string {
	l := len(s)
	if l == 0 {
		return ""
	}
	// 中心
	center := 0
	// 右边界
	rightBoundary := 0
	// 填充数据
	bs := []byte{'#'}
	for _, c := range s {
		bs = append(bs, byte(c))
		bs = append(bs, '#')
	}
	l = l*2 + 1
	// 半径数组
	ps := make([]int, l)
	count := true
	for i := range bs {
		count = true
		if i < rightBoundary {
			ps[i] = ps[2*center-i]
			if i+ps[i] > rightBoundary {
				ps[i] = rightBoundary - i
			}
			if i+ps[i] < rightBoundary {
				count = false
			}
		}

		if !count {
			continue
		}

		// 计算回文新长度
		for i+ps[i]+1 < l && i-ps[i]-1 >= 0 && bs[i+ps[i]+1] == bs[i-ps[i]-1] {
			ps[i]++
		}
		// 更新中心点
		if ps[i] > ps[center] {
			center = i
			rightBoundary = i + ps[i]
		}
	}

	// 过滤#
	buffer := &bytes.Buffer{}
	for i := center - ps[center] + 1; i < rightBoundary; i += 2 {
		buffer.WriteByte(bs[i])
	}

	return buffer.String()
}
