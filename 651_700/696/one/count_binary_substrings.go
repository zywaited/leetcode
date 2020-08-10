package one

func CountBinarySubstrings(s string) int {
	on := 0
	zn := 0
	pb := s[0]
	if s[0] == '0' {
		zn++
	} else {
		on++
	}
	ans := 0
	for i := 1; i < len(s); i++ {
		if s[i] == '0' {
			if pb != s[i] {
				zn = 0
			}
			pb = s[i]
			zn++
			if zn <= on {
				ans++
				continue
			}
			on = 0
			continue
		}
		if pb != s[i] {
			on = 0
		}

		pb = s[i]
		on++
		if on <= zn {
			ans++
			continue
		}
		zn = 0
	}
	return ans
}
