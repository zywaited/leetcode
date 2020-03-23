package one

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func Exist(board [][]byte, word string) bool {
	visited := make(map[int]byte)
	var dfs func(int, int, int) bool
	dfs = func(row int, col int, index int) bool {
		if index >= len(word) {
			return true
		}
		if row < 0 || row >= len(board) || col < 0 || col >= len(board[row]) {
			return false
		}
		if board[row][col] != word[index] {
			return false
		}
		if visited[row*201+col] == 1 {
			return false
		}
		visited[row*201+col] = 1
		for _, direct := range directs {
			if dfs(row+direct[0], col+direct[1], index+1) {
				return true
			}
		}
		visited[row*201+col] = 0
		return false
	}
	for row := range board {
		for col := range board[row] {
			if dfs(row, col, 0) {
				return true
			}
		}
	}
	return false
}
