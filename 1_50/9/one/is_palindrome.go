package one

func IsPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	rv := 0
	tmp := x
	for tmp != 0 {
		rv = tmp%10 + rv*10
		tmp /= 10
	}

	return rv == x
}
