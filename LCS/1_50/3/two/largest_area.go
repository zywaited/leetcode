package two

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func LargestArea(grid []string) int {
	m, n := len(grid), len(grid[0])
	indexes := make([]int, m*n+1)
	nums := make([]int, len(indexes))
	for index := range indexes {
		indexes[index] = index
		nums[index] = 1
	}
	find := func(index int) int {
		for indexes[index] != index {
			indexes[index] = indexes[indexes[index]]
			index = indexes[index]
		}
		return index
	}
	union := func(firstIndex, secondIndex int) {
		f := find(firstIndex)
		s := find(secondIndex)
		if indexes[f] == indexes[s] {
			return
		}
		if f < s {
			indexes[f] = indexes[s]
			nums[s] += nums[f]
			return
		}
		indexes[s] = indexes[f]
		nums[f] += nums[s]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 || i == m-1 || j == n-1 || grid[i][j] == '0' {
				union(i*n+j, len(indexes)-1)
				continue
			}
			for _, direct := range directs {
				ni, nj := i+direct[0], j+direct[1]
				if grid[ni][nj] == '0' {
					union(i*n+j, len(indexes)-1)
					continue
				}
				if grid[i][j] == grid[ni][nj] {
					union(i*n+j, ni*n+nj)
				}
			}
		}
	}
	ans := 0
	for index := range indexes {
		s := find(index)
		if s < len(indexes)-1 && nums[s] > ans {
			ans = nums[s]
		}
	}
	return ans
}
