package one

type SudoKu struct {
	// 长度为10，因为数字最大为9，索引也为9，不想数字减一操作
	// rows
	rows [9][10]int
	// cols
	cols [9][10]int
	// boxes
	boxes [9][10]int
	// 数独数据
	boards [][]byte
	// 是否结束(题目假设是唯一解)
	over bool
}

// 加入
func (sk *SudoKu) addSudoku(d, row, col int) {
	// 方块索引, 刚好分成9块
	idx := (row/3)*3 + col/3
	sk.rows[row][d]++
	sk.cols[col][d]++
	sk.boxes[idx][d]++
}

// 落子
func (sk *SudoKu) placeNumber(d, row, col int) {
	sk.addSudoku(d, row, col)
	// 放入数独中
	sk.boards[row][col] = byte(d) + '0'
}

// 删除落子
func (sk *SudoKu) removeNumber(d, row, col int) {
	// 方块索引
	idx := (row/3)*3 + col/3
	sk.rows[row][d]--
	sk.cols[col][d]--
	sk.boxes[idx][d]--
	// 还原空格
	sk.boards[row][col] = '.'
}

// 放置下一个子
func (sk *SudoKu) placeNextNumber(row, col int) {
	// 判断是否已经到结尾了
	if row == 8 && col == 8 {
		sk.over = true
		return
	}
	// 下一个，如果是下一行
	if col == 8 {
		sk.backtrace(row+1, 0)
		return
	}
	sk.backtrace(row, col+1)
}

// 判断是否可以落子，即当前行列方块都不存在当前子
func (sk *SudoKu) couldPlaceNumber(d, row, col int) bool {
	idx := (row/3)*3 + col/3
	return sk.rows[row][d]+sk.cols[col][d]+sk.boxes[idx][d] == 0
}

// 回溯
func (sk *SudoKu) backtrace(row, col int) {
	// 不是空格直接进入下一子
	if sk.boards[row][col] != '.' {
		sk.placeNextNumber(row, col)
		return
	}
	// 从1-9落子
	for d := 1; d < 10; d++ {
		// 是否存在
		if !sk.couldPlaceNumber(d, row, col) {
			continue
		}
		// 落子
		sk.placeNumber(d, row, col)
		// 下一子
		sk.placeNextNumber(row, col)
		// 如果结束
		if sk.over {
			return
		}
		// 递归调用结束未完成则删除当前子继续下
		sk.removeNumber(d, row, col)
	}
}

func SolveSudoku(board [][]byte) {
	sk := &SudoKu{boards: board, over: false}
	// 先初始化已经有的子
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			sk.addSudoku(int(board[i][j]-'0'), i, j)
		}
	}
	sk.backtrace(0, 0)
}
