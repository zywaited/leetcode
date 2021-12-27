package one

const mod = 1e9 + 7

var directs = [][]int{{0, -1}, {-1, 0}, {0, 1}, {1, 0}}

func FindPaths(m int, n int, maxMove int, startRow int, startColumn int) int {
	tm := m + 1
	tn := n + 1
	dp := make([][][]int, maxMove+1)
	for i := range dp {
		dp[i] = make([][]int, tm+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, tn+1)
		}
	}
	dp[0][startRow+1][startColumn+1] = 1
	ans := 0
	for mv := 1; mv <= maxMove; mv++ {
		for i := 0; i <= tm; i++ {
			for j := 0; j <= tn; j++ {
				for _, direct := range directs {
					ni, nj := i+direct[0], j+direct[1]
					if ni <= 0 || nj <= 0 || ni >= tm || nj >= tn || dp[mv-1][ni][nj] == 0 {
						continue
					}
					if dp[mv][i][j] == -1 {
						dp[mv][i][j] = 0
					}
					dp[mv][i][j] = (dp[mv][i][j] + dp[mv-1][ni][nj]) % mod
					if i == 0 || i == tm || j == 0 || j == tn {
						ans = (ans + dp[mv][i][j]) % mod
					}
				}
			}
		}
	}
	return ans
}
