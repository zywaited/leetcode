package two

func SmallestRepunitDivByK(k int) int {
	e := k % 10
	if e != 1 && e != 3 && e != 7 && e != 9 {
		return -1
	}
	l := 1
	n := 1
	for {
		n %= k
		if n == 0 {
			return l
		}
		l++
		n = n*10 + 1
	}
}
