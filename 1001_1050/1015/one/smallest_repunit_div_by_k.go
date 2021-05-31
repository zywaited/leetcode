package one

func SmallestRepunitDivByK(k int) int {
	n := 1
	nm := map[int]bool{}
	l := 1
	for {
		n %= k
		if n == 0 {
			return l
		}
		if nm[n] {
			return -1
		}
		nm[n] = true
		l++
		n = n*10 + 1
	}
}
