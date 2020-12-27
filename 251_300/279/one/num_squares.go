package one

import "math"

func NumSquares(n int) int {
	ns := make(map[int]int)
	var bc func(int, int) int
	bc = func(num int, ap int) int {
		if num == 0 || ns[num] > 0 {
			return ns[num]
		}
		mn := num
		ms := 0
		for ; ap > 0; ap-- {
			ms = ap * ap
			if ms > num {
				continue
			}
			if ms == num {
				mn = 1
				break
			}
			mn = min(mn, bc(num%ms, ap-1)+num/ms)
		}
		ns[num] = mn
		return mn
	}
	return bc(n, int(math.Sqrt(float64(n))))
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
