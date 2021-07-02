package one

var cbs = []string{"A", "C", "G", "T"}

func MinMutation(start string, end string, bank []string) int {
	bankM := make(map[string]bool, len(bank))
	for _, cb := range bank {
		bankM[cb] = true
	}
	type node struct {
		c string
		s int
		o int
	}
	vm := make(map[string]*node)
	q1 := []node{{c: start, o: 1}}
	q2 := []node{{c: end, o: 2}}
	vm[q1[0].c] = &q1[0]
	vm[q2[0].c] = &q2[0]
	for len(q1) > 0 && len(q2) > 0 {
		if len(q1) > len(q2) {
			q1, q2 = q2, q1
		}
		cn := q1[0]
		q1 = q1[1:]
		for i := range cn.c {
			for _, ncb := range cbs {
				if cn.c[i] == ncb[0] {
					continue
				}
				nc := cn.c[:i] + ncb + cn.c[i+1:]
				if !bankM[nc] {
					continue
				}
				if vm[nc] == nil {
					nn := node{c: nc, s: cn.s + 1, o: cn.o}
					vm[nc] = &nn
					q1 = append(q1, nn)
					continue
				}
				if vm[nc].o != cn.o {
					return cn.s + vm[nc].s + 1
				}
			}
		}
	}
	return -1
}
