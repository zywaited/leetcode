package one

func PushDominoes(dominoes string) string {
	bs := []byte(dominoes)
	for i := 0; i < len(dominoes); i++ {
		j := i
		for ; i < len(dominoes) && dominoes[i] == '.'; i++ {
		}
		if j == i {
			continue
		}
		p := byte('.')
		if j-1 >= 0 {
			p = bs[j-1]
		}
		c := byte('.')
		if i < len(dominoes) {
			c = dominoes[i]
		}
		if p != 'R' && c != 'L' {
			continue
		}
		k := i - 1
		for j < k {
			if p == 'R' {
				bs[j] = p
				j++
			}
			if c == 'L' {
				bs[k] = c
				k--
			}
		}
		if j != k {
			continue
		}
		if p == 'R' && c != 'L' {
			bs[j] = p
			continue
		}
		if c == 'L' && p != 'R' {
			bs[j] = c
		}
	}
	return string(bs)
}
