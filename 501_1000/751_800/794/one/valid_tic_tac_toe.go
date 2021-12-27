package one

func ValidTicTacToe(board []string) bool {
	x := 0
	o := 0
	xc := 0
	oc := 0
	c := 0
	for i := range board {
		for j := range board[i] {
			switch board[i][j] {
			case 'X':
				x++
			case 'O':
				o++
			}
			if board[i][j] == ' ' {
				continue
			}
			c = 0
			if j == 0 && board[i][j] == board[i][j+1] && board[i][j] == board[i][j+2] {
				c++
			}
			if i == 0 && board[i][j] == board[i+1][j] && board[i][j] == board[i+2][j] {
				c++
			}
			if i == 0 && j == 0 && board[i][j] == board[i+1][j+1] && board[i][j] == board[i+2][j+2] {
				c++
			}
			if i == len(board)-1 && j == 0 && board[i][j] == board[i-1][j+1] && board[i][j] == board[i-2][j+2] {
				c++
			}
			switch board[i][j] {
			case 'X':
				xc += c
			case 'O':
				oc += c
			}
		}
	}
	return (x-o == 1 && oc == 0 && xc <= 2) || (x == o && xc == 0 && oc <= 1)
}
