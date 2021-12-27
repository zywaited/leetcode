package one

var directs = [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func NumRookCaptures(board [][]byte) int {
	// 先找到R
	ri, rj := -1, -1
	for i := range board {
		if ri >= 0 && rj >= 0 {
			break
		}
		for j := range board {
			if board[i][j] == 'R' {
				ri = i
				rj = j
				break
			}
		}
	}
	// 然后4个方向移动
	ans := 0
	for _, direct := range directs {
		for step := 1; ; step++ {
			i, j := ri+direct[0]*step, rj+direct[1]*step
			if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) {
				break
			}
			if board[i][j] == 'B' {
				break
			}
			if board[i][j] == 'p' {
				ans++
				break
			}
		}
	}
	return ans
}
