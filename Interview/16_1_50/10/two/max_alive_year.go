package two

func MaxAliveYear(birth []int, death []int) int {
	ns := make([]int, 101) // 1900 - 2000 总共101年
	for _, y := range birth {
		ns[y-1900]++
	}
	for _, y := range death {
		if y < 2000 {
			// 减人口算在下一个年
			ns[y-1900+1]--
		}
	}
	sum := 0
	max := 0
	year := 0
	for i, n := range ns {
		sum += n
		if sum > max {
			max = sum
			year = i + 1900
		}
	}
	return year
}
