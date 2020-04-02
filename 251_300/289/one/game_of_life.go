package one

var directs = [][]int{{0, -1}, {-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}}

func GameOfLife(board [][]int) {
	for i := range board {
		for j := range board[i] {
			// 活细胞数量
			num := 0
			for _, direct := range directs {
				ni, nj := i+direct[0], j+direct[1]
				if ni < 0 || ni >= len(board) || nj < 0 || nj >= len(board[i]) {
					continue
				}
				if board[ni][nj]&1 == 1 {
					num++
				}
			}
			if board[i][j] == 1 {
				if num >= 2 && num <= 3 {
					continue
				}
				board[i][j] |= 2 // 标记死亡
				continue
			}
			if num == 3 {
				board[i][j] |= 2 // 标记复活
			}
		}
	}
	// 更正实际数据
	for i := range board {
		for j := range board[i] {
			if board[i][j]&2 == 0 {
				continue
			}
			board[i][j] ^= 3 // 原本死亡变复活，活着变死亡
		}
	}
}
