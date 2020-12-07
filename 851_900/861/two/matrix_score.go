package two

func MatrixScore(A [][]int) int {
	// 通过第一种解法我们知道了为了值最大
	// 1: 第一列必须全部为1
	// 2: 其余列如果需要移动，则当前1个数必须小于0
	r := len(A)
	c := len(A[0])
	s := (1 << uint(c-1)) * r // 第一列全为1
	for j := 1; j < c; j++ {
		on := 0 // 1的个数
		for i := 0; i < r; i++ {
			if A[i][0] == 1 {
				// 没有行移动
				on += A[i][j]
				continue
			}
			// 行移动后0<=>1
			on += A[i][j] ^ 1
		}
		if on < r-on {
			// 列移动
			on = r - on
		}
		s += (1 << uint(c-1-j)) * on
	}
	return s
}
