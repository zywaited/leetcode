package one

func SmallestNumber(pattern string) string {
	bs := make([]byte, 0, 9)
	n := 0
	for p := 0; p <= len(pattern); p++ {
		if p == len(pattern) || pattern[p] == 'I' {
			bs = append(bs, byte(n+1)+'0')
			n++
			continue
		}
		dn := 1
		for p++; p < len(pattern) && pattern[p] == 'D'; p++ {
			dn++
		}
		n += dn + 1
		for cn := 0; cn <= dn; cn++ {
			bs = append(bs, byte(n-cn)+'0')
		}
	}
	return string(bs)
}
