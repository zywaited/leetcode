package one

func RemoveInvalidParentheses(s string) []string {
	return newRemover(s).exec()
}

type remover struct {
	s   string
	lc  int
	rc  int
	lr  int
	rr  int
	ans []string
}

func newRemover(s string) *remover {
	r := remover{s: s}
	lc := 0
	for i := range s {
		if s[i] == '(' {
			r.lc++
			lc++
			continue
		}
		if s[i] != ')' {
			continue
		}
		r.rc++
		if lc > 0 {
			lc--
			continue
		}
		r.rr++
	}
	r.lr = lc
	return &r
}

func (r *remover) exec() []string {
	r.dfs(r.s, 0, 0, 0)
	return r.ans
}

func (r *remover) dfs(s string, pi, lr, rr int) {
	if lr == r.lr && rr == r.rr {
		if r.valid(s) {
			r.ans = append(r.ans, s)
		}
		return
	}
	for i := pi; i < len(s); i++ {
		if i != pi && s[i-1] == s[i] {
			continue
		}
		if len(s)-i < r.lr-lr+r.rr-rr {
			return
		}
		if r.lr-lr > 0 && s[i] == '(' {
			r.dfs(s[:i]+s[i+1:], i, lr+1, rr)
		}
		if r.rr-rr > 0 && s[i] == ')' {
			r.dfs(s[:i]+s[i+1:], i, lr, rr+1)
		}
	}
}

func (r *remover) valid(s string) bool {
	lc := 0
	for i := range s {
		if s[i] == '(' {
			lc++
			continue
		}
		if s[i] != ')' {
			continue
		}
		if lc > 0 {
			lc--
			continue
		}
		return false
	}
	return lc == 0
}
