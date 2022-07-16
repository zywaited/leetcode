package one

import "strings"

type WordFilter struct {
	wm  map[string]int
	psm map[string]int
}

func Constructor(words []string) WordFilter {
	wf := WordFilter{
		wm:  map[string]int{},
		psm: map[string]int{},
	}
	for index, word := range words {
		wf.wm[word] = index + 1
		for i := 1; i < len(word); i++ {
			for j := i + 1; j < len(word); j++ {
				wf.psm[word[:i]+":"+word[j:]] = index + 1
			}
		}
	}
	return wf
}

func (wf *WordFilter) F(pref string, suff string) int {
	index := max(wf.wm[pref+suff], wf.psm[pref+":"+suff])
	for i := 0; i < len(pref); i++ {
		if !strings.HasPrefix(suff, pref[i:]) {
			continue
		}
		index = max(index, wf.wm[pref+suff[len(pref[i:]):]])
	}
	return index - 1
}

func max(f, s int) int {
	if s <= f {
		return f
	}
	return s
}
