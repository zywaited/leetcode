package one

// IV = V-I, IX = X-I, XL = L-X, XC = C-X, CD = D-C, CM = M-C
var mp = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func RomanToInt(s string) int {
	pn := 0 // 记录前一个数字(用于计算IV这种结构)
	ans := 0
	for i := range s {
		num := mp[s[i]]
		ans += num
		if num > pn {
			ans -= pn << 1 // 去除多加的数据
		}
		pn = num
	}
	return ans
}
