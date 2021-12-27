package one

func BuddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	bn := [26]bool{}
	pi := -1
	pf := false
	for i := range s {
		if s[i] == goal[i] {
			pf = pf || bn[s[i]-'a']
			bn[s[i]-'a'] = true
			continue
		}
		if pi == len(s) {
			return false
		}
		if pi == -1 {
			pi = i
			continue
		}
		if s[pi] != goal[i] || s[i] != goal[pi] {
			return false
		}
		pi = len(s)
	}
	return pi == len(s) || (pi == -1 && pf)
}
