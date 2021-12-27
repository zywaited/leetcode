package one

func LongestPalindrome(s string) string {
	start := 0
	end := 0
	l := len(s)
	add := func(j, k int) int {
		for j >= 0 && k < l {
			if s[k] != s[j] {
				break
			}

			k++
			j--
		}

		return k - j - 1
	}

	even := 0
	odd := 0
	max := 0
	for i := range s {
		// 回文串可能是偶数也可能是奇数
		// 其实可以进行预处理成奇数，比如在每个字符中间加#，下个算法这样做
		even = add(i, i+1)
		odd = add(i-1, i+1)
		if even == 0 && odd == 0 {
			continue
		}

		max = even
		if odd > max {
			max = odd
		}

		if max > end-start {
			start = i - (max-1)/2
			end = i + max/2 + 1
		}
	}

	return s[start:end]
}
