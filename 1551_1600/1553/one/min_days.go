package one

func MinDays(n int) int {
	dm := make(map[int]int)
	var dfs func(int) int
	dfs = func(num int) int {
		if num <= 0 {
			return num
		}
		if dm[num] == 0 {
			two := num % 2
			if num/2 > 0 {
				two += dfs(num/2) + 1
			}
			three := num % 3
			if num/3 > 0 {
				three += dfs(num/3) + 1
			}
			dm[num] = min(two, three)
		}
		return dm[num]
	}
	return dfs(n)
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
