package one

func LastRemaining(n int, m int) int {
	num := 0
	for i := 2; i <= n; i++ {
		num = (m + num) % i
	}
	return num
}
