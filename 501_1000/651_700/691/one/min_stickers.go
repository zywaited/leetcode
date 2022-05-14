package one

import "math"

func MinStickers(stickers []string, target string) int {
	need := [26]int{}
	for i := range target {
		need[target[i]-'a']++
	}
	nums := make(map[[26]int]int)
	var bc func(i int, need [26]int) int
	bc = func(i int, need [26]int) int {
		if i == len(stickers) {
			return -1
		}
		if num, ok := nums[need]; ok {
			return num
		}
		fn := -1
		for ; i < len(stickers); i++ {
			have := [26]int{}
			for j := range stickers[i] {
				have[stickers[i][j]-'a']++
			}
			num := 0
			for j := range need {
				if need[j] > 0 && have[j] > 0 {
					num = max(num, int(math.Ceil(float64(need[j])/float64(have[j]))))
				}
			}
			nn := -1
			for ; num > 0; num-- {
				curr := need
				f := false
				for j := range have {
					curr[j] -= min(curr[j], have[j]*num)
					f = f || curr[j] > 0
				}
				cn := 0
				if f {
					cn = bc(i+1, curr)
				}
				if cn < 0 {
					continue
				}
				if nn < 0 {
					nn = num + cn
					continue
				}
				nn = min(nn, num+cn)
			}
			if nn < 0 {
				continue
			}
			if fn < 0 {
				fn = nn
				continue
			}
			fn = min(fn, nn)
		}
		nums[need] = fn
		return fn
	}
	return bc(0, need)
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
