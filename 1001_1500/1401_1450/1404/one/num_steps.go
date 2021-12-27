package one

func NumSteps(s string) int {
	ans := 0
	index := len(s) - 1
	for ; index > 0 || (index == 0 && s[index] == '0'); index-- {
		if s[index] == '0' {
			// >> 1
			ans++
			continue
		}
		// + 1 && >> 1
		ans += 2
		for index--; index >= 0; index-- {
			if s[index] == '0' {
				// == 0 时刚好是1
				if index > 0 {
					ans += 2
				}
				continue
			}
			ans++
		}
	}
	return ans
}
