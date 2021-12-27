package one

const mod = 1e9 + 7

func CountPairs(deliciousness []int) int {
	dm := make(map[int]int)
	ans := 0
	for _, num := range deliciousness {
		for i := uint(0); i <= 21; i++ {
			if 1<<i >= num {
				ans = (ans + dm[(1<<i)-num]) % mod
			}
		}
		dm[num]++
	}
	return ans
}
