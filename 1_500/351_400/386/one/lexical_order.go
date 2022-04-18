package one

func LexicalOrder(n int) []int {
	ans := make([]int, 0, n)
	cn := 1
	for cn <= n {
		for cn <= n {
			ans = append(ans, cn)
			cn *= 10
		}
		for cn = cn / 10; cn+1 <= n && cn%10+1 <= 9; {
			cn++
			ans = append(ans, cn)
		}
		for cn /= 10; cn%10 == 9; cn /= 10 {
		}
		cn++
		if cn == 1 {
			break
		}
	}
	return ans
}
