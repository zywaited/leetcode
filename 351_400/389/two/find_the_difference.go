package two

/**
 * map
 */
func FindTheDifference(s string, t string) byte {
	mp := make(map[byte]uint)
	i := 0
	sl := len(s)
	tl := len(t)
	for ; i < sl; i++ {
		mp[s[i]]++
	}
	for i = 0; i < tl; i++ {
		if mp[t[i]] != 0 {
			mp[t[i]]--
			continue
		}
		return t[i]
	}
	return 0
}
