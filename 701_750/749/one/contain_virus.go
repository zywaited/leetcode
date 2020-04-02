package one

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func ContainVirus(grid [][]int) int {
	ans := 0
	m, n := len(grid), len(grid[0])
	times := 1 // 处理次数
	visited := make(map[int]int, m*n)
	// 病毒区域
	queue := make([][]int, 0)
	// 索引对应的防火墙和感染数
	nums := make([][2]int, 0)
	zeroVisited := make(map[int]int, m*n)
	// 区域索引
	num := -1
	// 搜索
	var dfs func(int, int, int)
	dfs = func(sc, i, j int) {
		queue[num] = append(queue[num], i*51+j)
		visited[i*51+j] = times
		for _, direct := range directs {
			ni, nj := i+direct[0], j+direct[1]
			if ni < 0 || ni >= m || nj < 0 || nj >= len(grid[ni]) {
				continue
			}
			if grid[ni][nj] == 0 {
				if osc, ok := zeroVisited[ni*51+nj]; !ok || osc != sc {
					// 感染数+1
					nums[num][0]++
					zeroVisited[ni*51+nj] = sc
				}
				// 防火墙数+1
				nums[num][1]++
				continue
			}
			if grid[ni][nj] == 1 && visited[ni*51+nj] != times {
				dfs(sc, ni, nj)
			}
		}

	}
	for ; ; times++ {
		// 重置
		queue = queue[:0]
		nums = nums[:0]
		num = -1
		zeroVisited = make(map[int]int, m*n)
		mi := -1 // 需要安装防火墙的区域
		mn := 0
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] != 1 || visited[i*51+j] == times {
					continue
				}
				queue = append(queue, []int{})
				nums = append(nums, [2]int{})
				num++
				dfs(i*51+j, i, j)
				if mi == -1 || nums[num][0] > mn {
					mi = num
					mn = nums[num][0]
				}
			}
		}
		if mi == -1 || mn == 0 {
			// 一个都不会感染
			break
		}
		ans += nums[mi][1]
		for i := range queue {
			for _, sc := range queue[i] {
				if i == mi {
					grid[sc/51][sc%51] = -1 // 只要不等于0和1就行
					continue
				}
				// 对周围进行感染
				for _, direct := range directs {
					ni, nj := sc/51+direct[0], sc%51+direct[1]
					if ni < 0 || ni >= m || nj < 0 || nj >= len(grid[ni]) || grid[ni][nj] != 0 {
						continue
					}
					grid[ni][nj] = 1
				}
			}
		}
	}
	return ans
}
