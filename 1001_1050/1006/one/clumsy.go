package one

func Clumsy(n int) int {
	ans := 0
	cn := 0
	f := true
	for n > 0 {
		if n > 0 {
			cn = n
			n--
		}
		if n > 0 {
			cn *= n
			n--
		}
		if n > 0 {
			cn /= n
			n--
		}
		if n > 0 {
			if f {
				cn += n
			} else {
				cn -= n
			}
			n--
		}
		if f {
			ans = cn
			f = false
			continue
		}
		ans -= cn
	}
	return ans
}
