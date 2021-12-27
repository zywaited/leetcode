package two

func CanMakeArithmeticProgression(arr []int) bool {
	min, max := arr[0], arr[0]
	nm := make(map[int]bool, len(arr))
	for _, n := range arr {
		nm[n] = true
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	if (max-min)%(len(arr)-1) > 0 {
		return false
	}
	v := (max - min) / (len(arr) - 1)
	cn := min
	for i := 0; i < len(arr); i++ {
		if !nm[cn] {
			return false
		}
		cn += v
	}
	return true
}
