package one

func addSpaces(s string, spaces []int) string {
	bs := []byte{}
	j := 0
	for i := range s {
		if j == len(spaces) || i < spaces[j] {
			bs = append(bs, s[i])
			continue
		}
		bs = append(bs, ' ', s[i])
		j++
	}
	return string(bs)
}
