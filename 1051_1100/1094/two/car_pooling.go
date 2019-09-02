package two

func CarPooling(trips [][]int, capacity int) bool {
	lt := [1001]int{} // 坐标人数
	for i := 0; i < len(trips); i++ {
		t := trips[i]
		lt[t[1]] += t[0] // 该点上车人数
		lt[t[2]] -= t[0] // 该点下车人数
	}
	for i := 0; i <= 1000; i++ {
		capacity -= lt[i] // 如果已经下车，则座位增加
		if capacity < 0 {
			return false
		}
	}
	return true
}
