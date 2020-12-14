package one

func GroupAnagrams(strs []string) [][]string {
	ans := make([][]string, 1)
	im := make(map[[26]int]int)
	cns := [26]int{}
	for _, str := range strs {
		for c := 0; c < 26; c++ {
			cns[c] = 0 // reset
		}
		for ci := range str {
			cns[str[ci]-'a']++
		}
		if im[cns] > 0 {
			ans[im[cns]] = append(ans[im[cns]], str)
			continue
		}
		ans = append(ans, []string{str})
		im[cns] = len(ans) - 1
	}
	return ans[1:]
}
