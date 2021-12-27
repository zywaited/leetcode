package one

import "math"

func ReorderedPowerOf2(n int) bool {
	nums := [10]int{}
	nl := 0
	for ; n > 0; n /= 10 {
		nums[n%10]++
		nl++
	}
	mn := int(math.Pow10(nl - 1))
	cn := 1
	for ; cn/10 < mn; cn <<= 1 {
		if cn < mn {
			continue
		}
		tmp := [10]int{}
		for tn := cn; tn > 0; tn /= 10 {
			tmp[tn%10]++
		}
		f := true
		for i := range tmp {
			f = f && nums[i] == tmp[i]
		}
		if f {
			return true
		}
	}
	return false
}
