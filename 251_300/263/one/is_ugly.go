package one

func IsUgly(num int) bool {
	if num <= 0 {
		return false
	}
	for _, nf := range []int{2, 3, 5} {
		for num%nf == 0 {
			num /= nf
		}
	}
	return num == 1
}
