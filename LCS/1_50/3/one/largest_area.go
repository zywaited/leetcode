package one

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func LargestArea(grid []string) int {
	vm := make(map[int]int)
	ans := 0
	var bc func(i, j, n int) (nn int)
	bc = func(i, j, n int) (nn int) {
		vm[i*500+j] = 1
		defer func() {
			if nn == 0 {
				vm[i*500+j] = 2
			}
		}()
		for _, direct := range directs {
			ni, nj := i+direct[0], j+direct[1]
			if grid[ni][nj] == '0' {
				return 0
			}
			if grid[ni][nj] != grid[i][j] {
				continue
			}
			if ni == 0 || ni == len(grid)-1 || nj == 0 || nj == len(grid[ni])-1 {
				return 0
			}
			s := vm[ni*500+nj]
			if s == 1 {
				continue
			}
			if s == 2 {
				return 0
			}
			nn = bc(ni, nj, 1)
			if nn == 0 {
				return
			}
			n += nn
		}
		nn = n
		return
	}
	for i := 1; i < len(grid)-1; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			if vm[i*500+j] > 0 || grid[i][j] == '0' {
				continue
			}
			n := bc(i, j, 1)
			if n > ans {
				ans = n
			}
		}
	}
	return ans
}
