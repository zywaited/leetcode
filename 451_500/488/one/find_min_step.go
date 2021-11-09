package one

func FindMinStep(board string, hand string) int {
	return newStep().bfs(board, hand)
}

type step map[[2]string]int

func newStep() step {
	return make(step)
}

func (s step) bfs(board string, hand string) int {
	q := [][2]string{{board, hand}}
	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		t := s[c]
		for i := 0; i < len(c[0]); i++ {
			for j := 0; j < len(c[1]); j++ {
				cb := s.eliminate(c[0][:i] + c[1][j:j+1] + c[0][i:])
				if cb == "" {
					return t + 1
				}
				ch := c[1][:j] + c[1][j+1:]
				if ch == "" {
					continue
				}
				nc := [2]string{cb, ch}
				if _, ok := s[nc]; ok {
					continue
				}
				s[nc] = t + 1
				q = append(q, nc)
			}
		}
	}
	return -1
}

func (s step) eliminate(board string) string {
	bs := []byte{}
	q := [][2]int{}
	bl := 0
	num := 0
	for i := 0; i <= len(board); i++ {
		if i < len(board) && (i == 0 || board[i] == board[i-1]) {
			num++
			bs = append(bs, board[i])
			continue
		}
		if num >= 3 {
			bs = bs[:bl]
		}
		if num < 3 {
			q = append(q, [2]int{i - 1, num})
		}
		bl = len(bs)
		num = 1
		if i < len(board) {
			bs = append(bs, board[i])
			if len(q) > 0 && board[q[len(q)-1][0]] == board[i] {
				num += q[len(q)-1][1]
				bl -= q[len(q)-1][1]
				q = q[:len(q)-1]
			}
		}
	}
	return string(bs)
}
