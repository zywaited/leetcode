package one

func NthMagicalNumber(n int, a int, b int) int {
	mod := int(1e9 + 7)
	m := a * b / gcd(a, b)
	mn := m/a + m/b - 1

	ai := 1
	bi := 1
	an := 0
	for dn := n % mn; dn > 0; dn-- {
		an = min(a*ai, b*bi)
		if an == a*ai {
			ai++
		} else {
			bi++
		}
	}
	return (m*(n/mn) + an) % mod
}

func gcd(f, s int) int {
	for s > 0 {
		f, s = s, f%s
	}
	return f
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
