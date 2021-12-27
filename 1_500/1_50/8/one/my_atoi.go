package one

func MyAtoi(str string) int {
	l := len(str)
	if l < 1 {
		return 0
	}

	max := 1<<31 - 1
	min := -1 << 31
	mc, mp := max/10, max%10
	rv := 0
	negative := false
	first := false
	pop := 0
	for i := 0; i < l; i++ {
		if str[i] == ' ' {
			if first {
				break
			}

			continue
		}

		if str[i] == '-' || str[i] == '+' {
			if first {
				break
			}

			if str[i] == '-' {
				negative = true
			}

			first = true
			continue
		}

		if str[i] < '0' || str[i] > '9' {
			if !first {
				return 0
			}

			break
		}

		pop = int(str[i] - '0')
		// 边界值判断
		if negative && (rv > mc || (rv == mc && pop > mp)) {
			return min
		}

		if !negative && (rv > mc || (rv == mc && pop > mp-1)) {
			return max
		}

		rv = rv*10 + pop
		if !first {
			first = true
		}
	}

	if negative {
		rv = 0 - rv
	}

	return rv
}
