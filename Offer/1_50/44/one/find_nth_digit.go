package one

func FindNthDigit(n int) int {
	if n == 0 {
		return n
	}
	tn := 9
	sn := 1
	nn := 0
	for sn <= n {
		nn++
		sn += nn * tn
		tn *= 10
	}
	if nn == 1 {
		sn = 0
		tn = 10
		nn = 1
	}
	if nn > 1 {
		tn /= 10
		sn -= nn * tn
		tn /= 9
	}
	sn = n - sn
	fn := sn/(nn*tn) + 1
	en := sn % (nn * tn)
	if en == 0 {
		return fn
	}
	sn = fn*tn + en/nn
	en %= nn
	for en = nn - en; en > 0; en-- {
		fn = sn % 10
		sn /= 10
	}
	return fn
}
