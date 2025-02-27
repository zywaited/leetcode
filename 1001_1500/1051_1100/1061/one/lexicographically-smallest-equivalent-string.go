package one

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	const chars = 26
	parent := [chars]int{}
	for i := range parent {
		parent[i] = i
	}
	find := func(i int) int {
		for parent[i] != i {
			parent[i] = parent[parent[i]]
			i = parent[i]
		}
		return i
	}
	union := func(fi, si int) {
		f := find(fi)
		s := find(si)
		if f <= s {
			parent[s] = f
			return
		}
		parent[f] = s
	}
	for i := range s1 {
		union(int(s1[i]-'a'), int(s2[i]-'a'))
	}
	var bs []byte
	for i := range baseStr {
		bs = append(bs, byte(find(int(baseStr[i]-'a')))+'a')
	}
	return string(bs)
}
