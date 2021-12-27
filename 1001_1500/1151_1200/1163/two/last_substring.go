package two

// 最大表示法(最小表示法的一种变形)
func LastSubstring(s string) string {
	// 因为按照此方法是找到开头的字母
	// 而且是排列方式，所以需要添加一位比a小的字符
	// 例如"bcac"会找到"cac"和"cbc"，这时会返回"c"
	// 下面的一个方法(AnotherLastSubstring)加个条件也可以规避该问题
	s += "0"
	i, j, k := 0, 1, 0
	n := len(s)
	for i < n && j < n && k < n {
		ii := (i + k) % n
		ji := (j + k) % n
		if s[ii] == s[ji] {
			k++
			continue
		}
		if s[ii] > s[ji] {
			j += k + 1
		} else {
			i += k + 1
		}
		if i == j {
			j++
		}
		k = 0
	}
	min := i
	if j < min {
		min = j
	}
	return s[min : n-1]
}

// 针对上诉情况添加判断条件
func AnotherLastSubstring(s string) string {
	i, j, k := 0, 1, 0
	n := len(s)
	for i < n && j < n && k < n {
		// j会先到结尾
		// 这样直接退出就行了，肯定i开头的长度大
		if j+k >= n {
			break
		}
		if s[i+k] == s[j+k] {
			k++
			continue
		}
		if s[i+k] > s[j+k] {
			j += k + 1
		} else {
			i += k + 1
		}
		if i == j {
			j++
		}
		k = 0
	}
	return s[i:]
}
