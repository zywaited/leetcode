package one

func FindNthDigit(n int) int {
	base := 9
	cn := 9
	bn := 1
	for cn < n {
		base *= 10
		bn++
		cn += base * bn
	}
	nn := n - (cn - base*bn)
	pn := nn / bn
	en := nn % bn
	if en > 0 {
		pn++
	}
	num := base/9 + pn - 1
	for rn := bn - en; en > 0 && rn > 0; rn-- {
		num /= 10
	}
	return num % 10
}
