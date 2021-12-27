package one

type node struct {
	s string
	n int
	o int
}

func OpenLock(deadends []string, target string) int {
	s := "0000"
	if target == s {
		return 0
	}
	dm := make(map[string]bool, len(deadends))
	for _, deadend := range deadends {
		if deadend == s {
			return -1
		}
		dm[deadend] = true
	}
	q1 := []*node{{s: "0000"}}
	q2 := []*node{{s: target, o: 1}}
	vs := make(map[string]*node)
	vs[q1[0].s] = q1[0]
	vs[q2[0].s] = q2[0]
	for len(q1) > 0 || len(q2) > 0 {
		if len(q1) > 0 {
			cn := q1[0]
			q1 = q1[1:]
			for i := 0; i < len(cn.s); i++ {
				ns := cn.s[:i] + string(byte((int(cn.s[i]-'0')+9)%10)+'0') + cn.s[i+1:]
				ln := vs[ns]
				if ln == nil && !dm[ns] {
					nn := &node{s: ns, n: cn.n + 1}
					q1 = append(q1, nn)
					vs[ns] = nn
				}
				if ln != nil && ln.o == 1 {
					return cn.n + ln.n + 1
				}
				ns = cn.s[:i] + string(byte((int(cn.s[i]-'0')+1)%10)+'0') + cn.s[i+1:]
				ln = vs[ns]
				if ln == nil && !dm[ns] {
					nn := &node{s: ns, n: cn.n + 1}
					q1 = append(q1, nn)
					vs[ns] = nn
				}
				if ln != nil && ln.o == 1 {
					return cn.n + ln.n + 1
				}
			}
		}
		if len(q2) > 0 {
			cn := q2[0]
			q2 = q2[1:]
			for i := 0; i < len(cn.s); i++ {
				ns := cn.s[:i] + string(byte((int(cn.s[i]-'0')+9)%10)+'0') + cn.s[i+1:]
				ln := vs[ns]
				if ln == nil && !dm[ns] {
					nn := &node{s: ns, n: cn.n + 1, o: 1}
					q2 = append(q2, nn)
					vs[ns] = nn
				}
				if ln != nil && ln.o == 0 {
					return cn.n + ln.n + 1
				}
				ns = cn.s[:i] + string(byte((int(cn.s[i]-'0')+1)%10)+'0') + cn.s[i+1:]
				ln = vs[ns]
				if ln == nil && !dm[ns] {
					nn := &node{s: ns, n: cn.n + 1, o: 1}
					q2 = append(q2, nn)
					vs[ns] = nn
				}
				if ln != nil && ln.o == 0 {
					return cn.n + ln.n + 1
				}
			}
		}
	}
	return -1
}
