package one

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}, {1, 1}, {-1, -1}, {1, -1}, {-1, 1}}

func UpdateBoard(board [][]byte, click []int) [][]byte {
	queue := make([]int, 0, len(board)*len(board[0]))
	queue = append(queue, click[0]*51+click[1])
	ql := 0
	for len(queue) > 0 {
		r, c := queue[0]/51, queue[0]%51
		queue = queue[1:]
		switch board[r][c] {
		case 'M':
			board[r][c] = 'X'
		case 'E':
			num := 0
			ql = len(queue)
			for _, direct := range directs {
				nr, nc := r+direct[0], c+direct[1]
				if nr < 0 || nc < 0 || nr >= len(board) || nc >= len(board[nr]) {
					continue
				}
				switch board[nr][nc] {
				case 'E':
					queue = append(queue, nr*51+nc)
				case 'M', 'X':
					num++
				}
			}
			if num > 0 {
				board[r][c] = byte(num) + '0'
				queue = queue[:ql]
				continue
			}
			board[r][c] = 'B'
		}
	}
	return board
}
