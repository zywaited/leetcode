package one

func PlaceWordInCrossword(board [][]byte, word string) bool {
	cols := make([]int, len(board[0]))
	asc := func(i, j, num int) bool {
		if num != len(word) {
			return false
		}
		wi := 0
		for nj := j - num; nj < j; nj++ {
			if board[i][nj] == ' ' || board[i][nj] == word[wi] {
				wi++
				continue
			}
			return false
		}
		return true
	}
	desc := func(i, j, num int) bool {
		if num != len(word) {
			return false
		}
		wi := len(word) - 1
		for nj := j - num; nj < j; nj++ {
			if board[i][nj] == ' ' || board[i][nj] == word[wi] {
				wi--
				continue
			}
			return false
		}
		return true
	}
	colAsc := func(i, j, num int) bool {
		if num != len(word) {
			return false
		}
		wi := 0
		for ni := i - num; ni < i; ni++ {
			if board[ni][j] == ' ' || board[ni][j] == word[wi] {
				wi++
				continue
			}
			return false
		}
		return true
	}
	colDesc := func(i, j, num int) bool {
		if num != len(word) {
			return false
		}
		wi := len(word) - 1
		for ni := i - num; ni < i; ni++ {
			if board[ni][j] == ' ' || board[ni][j] == word[wi] {
				wi--
				continue
			}
			return false
		}
		return true
	}
	for i := range board {
		prev := 0
		for j := range board[i] {
			if board[i][j] != '#' {
				prev++
				cols[j]++
				continue
			}
			if asc(i, j, prev) || desc(i, j, prev) || colAsc(i, j, cols[j]) || colDesc(i, j, cols[j]) {
				return true
			}
			prev = 0
			cols[j] = 0
		}
		j := len(board[i])
		if asc(i, j, prev) || desc(i, j, prev) {
			return true
		}
	}
	for j := range cols {
		if colAsc(len(board), j, cols[j]) || colDesc(len(board), j, cols[j]) {
			return true
		}
	}
	return false
}
