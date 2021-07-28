package one

import "sort"

func MinDeletions(s string) int {
	cns := [26]int{}
	for i := range s {
		cns[s[i]-'a']++
	}
	sort.Ints(cns[:])
	pn := cns[len(cns)-1]
	ans := 0
	for index := len(cns) - 2; index >= 0 && cns[index] > 0; index-- {
		if cns[index] >= pn {
			ans += cns[index] - pn
			if pn > 0 {
				ans++
				pn--
			}
			continue
		}
		pn = cns[index]
	}
	return ans
}
