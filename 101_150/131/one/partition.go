package one

func Partition(s string) [][]string {
	ans := make([][]string, 0)
	temp := make([]string, 0)
	if len(s) == 0 {
		ans = append(ans, temp)
		return ans
	}
	// i-j是否是回文串
	pp := make([][]bool, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		pp[i] = make([]bool, len(s))
		for j := i; j < len(s); j++ {
			if s[i] == s[j] && (j-i <= 2 || pp[i+1][j-1]) {
				pp[i][j] = true
			}
		}
	}
	var bfs func(int)
	bfs = func(i int) {
		if i == len(s) {
			ans = append(ans, append(make([]string, 0, len(temp)), temp...))
			return
		}
		for j := i; j < len(s); j++ {
			if pp[i][j] {
				temp = append(temp, s[i:j+1])
				bfs(j + 1)
				temp = temp[:len(temp)-1]
			}
		}
	}
	bfs(0)
	return ans
}
