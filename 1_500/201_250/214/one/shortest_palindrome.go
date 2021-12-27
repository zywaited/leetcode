package one

import "bytes"

func reverse(s string) string {
	l := len(s)
	buffer := &bytes.Buffer{}
	for i := l - 1; i >= 0; i-- {
		buffer.WriteByte(s[i])
	}
	return buffer.String()
}

func ShortestPalindrome(s string) string {
	l := len(s)
	if l < 1 {
		return ""
	}
	str := reverse(s)
	tmp := s + "#" + str
	indexes := next(tmp)
	buffer := &bytes.Buffer{}
	index := indexes[len(tmp)-1]
	if index < 0 {
		index = 0
	}

	// 写入缺失的字符
	for i := l - 1; i > index; i-- {
		buffer.WriteByte(s[i])
	}
	// 写入原字符串
	buffer.WriteString(s)
	return buffer.String()
}

// 计算移动数组
func next(s string) []int {
	l := len(s)
	bs := make([]int, l)
	bs[0] = -1
	k := -1
	i := 0
	for i < l-1 {
		if k == -1 || s[i] == s[k] {
			k++
			i++
			bs[i] = k
			continue
		}

		k = bs[k]
	}

	return bs
}
