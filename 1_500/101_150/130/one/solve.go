package one

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func Solve(board [][]byte) {
	for i := range board {
		for j := range board[i] {
			if board[i][j] != 'O' {
				continue
			}
			if dfs(board, i, j) {
				change(board, i, j)
			}
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'T' {
				board[i][j] = 'O'
			}
		}
	}
}

func dfs(board [][]byte, i, j int) bool {
	if !valid(board, i, j) {
		return false
	}
	if board[i][j] != 'O' {
		return true
	}
	board[i][j] = 'T'
	flag := true
	for _, direct := range directs {
		ni, nj := i+direct[0], j+direct[1]
		if !dfs(board, ni, nj) {
			flag = false
		}
	}
	return flag
}

func change(board [][]byte, i, j int) {
	if !valid(board, i, j) || board[i][j] != 'T' {
		return
	}
	board[i][j] = 'X'
	for _, direct := range directs {
		change(board, i+direct[0], j+direct[1])
	}
}

func valid(board [][]byte, i, j int) bool {
	return i >= 0 && i < len(board) && j >= 0 && j < len(board[i])
}
