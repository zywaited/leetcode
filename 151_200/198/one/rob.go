package one

// f(n) = max(f(n-2)+An, f(n-1))
func Rob(nums []int) int {
	f := 0 // f(n-1)
	s := 0 // f(n-2)
	t := 0
	for _, num := range nums {
		t = f // 上一次的最大值
		if s+num > f {
			f = s + num
		}
		s = t // f(n-1) => f(n-2)
	}
	return f
}
