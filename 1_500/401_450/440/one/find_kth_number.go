package one

func FindKthNumber(n int, k int) int {
	ans := 1
	cn := 1
	tk := 0
	for tk < k {
		cnt, tn := count(cn, n)
		tk += cnt
		if tk < k {
			cn++
			continue
		}
		if tk == k {
			ans = tn
			break
		}
		ans = cn
		tk = tk - cnt + 1
		if cn*10 <= n {
			cn *= 10
			continue
		}
		cn++
	}
	return ans
}

func count(cn, n int) (int, int) {
	cnt := 0
	nnt := 1
	tn := cn
	for cn <= n {
		if cn*10 <= n {
			cnt += nnt
			tn = cn + nnt - 1
			cn *= 10
			nnt *= 10
			continue
		}
		if cn+nnt-1 <= n {
			cnt += nnt
			tn = cn + nnt - 1
			break
		}
		cnt += n - cn + 1
		tn = n
		break
	}
	return cnt, tn
}
