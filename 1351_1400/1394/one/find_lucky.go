package one

func FindLucky(arr []int) int {
	nm := make(map[int]int, len(arr))
	for _, num := range arr {
		nm[num]++
	}
	ans := -1
	for n, m := range nm {
		if n == m && n > ans {
			ans = n
		}
	}
	return ans
}
