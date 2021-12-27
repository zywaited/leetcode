package one

func GenerateParenthesis(n int) []string {
	rs := make([]string, 0)
	bc(nil, 0, 0, n, &rs)
	return rs
}

func bc(cur []byte, open, close, n int, rs *[]string) {
	if open == n && close == n {
		*rs = append(*rs, string(cur))
		return
	}
	// 左括号
	if open < n {
		bc(append(cur, '('), open+1, close, n, rs)
	}
	// 右括号
	if close < open && close < n {
		bc(append(cur, ')'), open, close+1, n, rs)
	}
}
