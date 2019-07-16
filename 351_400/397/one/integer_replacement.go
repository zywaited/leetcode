package one

/**
 * 规律
 * 3为分界线(本身除外)，这种奇数需要先加1再除以2次数是最小的
 */

func IntegerReplacement(n int) int {
	r := 0
	for n != 1 {
		r++
		// 等于3为加2
		if n == 3 {
			r++
			break
		}
		// 与操作，这类奇数先加1是最小
		if n&3 == 3 {
			n++
			continue
		}
		if n&1 == 1 {
			n--
			continue
		}
		n >>= 1
	}
	return r
}
