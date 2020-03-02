package one

func SumNums(n int) int {
	var (
		sum int
		add func(int) bool
	)
	add = func(num int) bool {
		sum += num
		return num > 1 && add(num)
	}
	add(n)
	return sum
}
