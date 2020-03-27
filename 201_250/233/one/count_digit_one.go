package one

import "math"

func CountDigitOne(n int) int {
	// 记录每个位置的数字
	nn := make([]int, 0)
	ans := 0
	tn := n
	pn := 0
	for tn > 0 {
		// 获取当前的数字
		pn = tn % 10
		tn /= 10
		nn = append(nn, pn)
		if tn == 0 {
			break
		}
		// 这个时候位数还有多的
		ans += int(math.Pow10(len(nn) - 1))
		if len(nn)-2 >= 0 {
			ans += 9 * int(math.Pow10(len(nn)-2)) * (len(nn) - 1)
		}
	}
	// 依次变化1的位置
	for i := len(nn) - 1; i >= 0; i-- {
		if nn[i] > 1 {
			ans += int(math.Pow10(i))
		} else if nn[i] == 1 {
			ans++ // 本身就是1，所以这里+1防止漏算
			// == 1情况比较复杂，需要(i-1)-0都计算一次
			for k := i - 1; k >= 0; k-- {
				if nn[k] > 0 {
					ans += nn[k] * int(math.Pow10(k))
				}
			}
		}
		for j := i + 1; j < len(nn); j++ {
			if j == len(nn)-1 {
				if nn[j] > 1 {
					ans += (nn[j] - 1) * int(math.Pow10(j-1))
				}
				continue
			}
			if nn[j] > 0 {
				ans += nn[j] * int(math.Pow10(j-1))
			}
		}
	}
	return ans
}
