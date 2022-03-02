package one

import (
	"math"
	"strconv"
)

func NearestPalindromic(n string) string {
	num, _ := strconv.ParseInt(n, 10, 64)
	if len(n) == 1 {
		return strconv.FormatInt(num-1, 10)
	}
	// 1: 新增一位 100....001
	f := int64(math.Pow(float64(10), float64(len(n)))) + 1
	// 2: 减少一位 999....999
	s := int64(math.Pow(float64(10), float64(len(n)-1))) - 1

	m := f
	if num-s <= f-num {
		m = s
	}

	bs := []byte(n)
	i := 0
	j := len(n) - 1
	for i < j {
		if bs[i] < '9' {
			// 3: 首位加1
			c := int64(0)
			for k := len(bs) - 1; k >= 0; k-- {
				t := int64(bs[k] - '0')
				if k == i {
					t++
				}
				if k == j {
					t = int64(bs[i] - '0' + 1)
				}
				if i < k && k < j {
					t = 0
				}
				c = c*10 + t
			}
			if (abs(c-num) < abs(m-num)) || (abs(c-num) == abs(m-num) && c < m) {
				m = c
			}
		}
		if bs[i] > '0' {
			// 4: 首位减1
			c := int64(0)
			for k := len(bs) - 1; k >= 0; k-- {
				t := int64(bs[k] - '0')
				if k == i {
					t--
				}
				if k == j {
					t = int64(bs[i] - '0' - 1)
				}
				if i < k && k < j {
					t = 9
				}
				c = c*10 + t
			}
			if (abs(c-num) < abs(m-num)) || (abs(c-num) == abs(m-num) && c < m) {
				m = c
			}
		}
		// 5: 首位不变
		bs[j] = bs[i]
		i++
		j--
	}
	// 与原来数字不同
	c, _ := strconv.ParseInt(string(bs), 10, 64)
	if c != num && ((abs(c-num) < abs(m-num)) || (abs(c-num) == abs(m-num) && c < m)) {
		m = c
	}
	// 中间数字可以变换
	if i == j {
		if bs[i] > '0' {
			bs[i]--
			c, _ := strconv.ParseInt(string(bs), 10, 64)
			if (abs(c-num) < abs(m-num)) || (abs(c-num) == abs(m-num) && c < m) {
				m = c
			}
			bs[i]++
		}
		if bs[i] < '9' {
			bs[i]++
			c, _ := strconv.ParseInt(string(bs), 10, 64)
			if (abs(c-num) < abs(m-num)) || (abs(c-num) == abs(m-num) && c < m) {
				m = c
			}
		}
	}
	return strconv.FormatInt(m, 10)
}

func abs(f int64) int64 {
	if f <= 0 {
		return -f
	}
	return f
}
