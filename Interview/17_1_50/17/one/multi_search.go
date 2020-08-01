package one

func MultiSearch(big string, smalls []string) [][]int {
	bm := make([][]int, 26)
	for index := range big {
		bm[big[index]-'a'] = append(bm[big[index]-'a'], index)
	}
	ans := make([][]int, len(smalls))
	for index, small := range smalls {
		if len(small) == 0 || len(bm[small[0]-'a']) == 0 {
			continue
		}
		for _, bi := range bm[small[0]-'a'] {
			if bi+len(small) > len(big) {
				break
			}
			if big[bi:bi+len(small)] == small {
				ans[index] = append(ans[index], bi)
			}
		}
	}
	return ans
}
