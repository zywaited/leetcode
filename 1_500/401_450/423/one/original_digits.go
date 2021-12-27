package one

func OriginalDigits(s string) string {
	bn := [26]int{}
	for i := range s {
		bn[s[i]-'a']++
	}
	nn := [10]int{}
	nn[0] = bn['z'-'a']
	nn[2] = bn['w'-'a']
	nn[4] = bn['u'-'a']
	nn[6] = bn['x'-'a']
	nn[8] = bn['g'-'a']
	nn[1] = bn['o'-'a'] - nn[0] - nn[2] - nn[4]
	nn[3] = bn['t'-'a'] - nn[2] - nn[8]
	nn[5] = bn['f'-'a'] - nn[4]
	nn[7] = bn['s'-'a'] - nn[6]
	nn[9] = bn['i'-'a'] - nn[5] - nn[8] - nn[6]

	bs := make([]byte, 0)
	for n := 0; n <= 9; n++ {
		for t := nn[n]; t > 0; t-- {
			bs = append(bs, byte(n)+'0')
		}
	}
	return string(bs)
}
