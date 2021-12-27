package one

// 最开始的想法是计算所有的blocked是否能够围住两个节点，这样就不能访问到
// 但是边境计算太复杂，节点要到四周
// 后面看到200点节点最大可以包住(200-1)*200/2个网格，这样就不用算围住的区域，就按最大网格判断

const (
	maxBlockedPointerNum = (200 - 1) * 200 / 2
	whiteTag             = 0 // 空白
	blockTag             = 1 // block网格
	sourceTag            = 2 // 起点
	targetTag            = 3 // 终点

	directNum = 4 // 方向
	min       = -1
	max       = 1e6
)

// 上下左右方向
var (
	dx = []int{-1, 1, 0, 0}
	dy = []int{0, 0, -1, 1}
)

type pointerStatus map[int]map[int]int

func (ps pointerStatus) add(pointer []int, tag int) {
	if ps[pointer[0]] == nil {
		ps[pointer[0]] = make(map[int]int)
	}
	ps[pointer[0]][pointer[1]] = tag
}

func (ps pointerStatus) get(pointer []int) int {
	if ps[pointer[0]] != nil {
		return ps[pointer[0]][pointer[1]]
	}
	return whiteTag
}

func IsEscapePossible(blocked [][]int, source []int, target []int) bool {
	if len(blocked) == 0 {
		return true
	}
	blockPointers := make(pointerStatus)
	for _, bk := range blocked {
		blockPointers.add(bk, blockTag)
	}
	blockPointers.add(source, sourceTag)
	blockPointers.add(target, targetTag)
	// 因为是根据最大网格数判断, 所以需要两边都要判断
	return bfs(blockPointers, source, sourceTag, targetTag) && bfs(blockPointers, target, targetTag, sourceTag)
}

// 广度搜索查看是否能够到终点
func bfs(blockPointers pointerStatus, source []int, sourceTag, targetTag int) bool {
	queue := [][]int{source}
	totalPointerNum := 0 // 包住的网格数量
	for len(queue) > 0 {
		pi := queue[0]
		queue = queue[1:]
		if blockPointers.get(pi) == targetTag {
			// 终点
			return true
		}
		totalPointerNum++
		if totalPointerNum > maxBlockedPointerNum {
			// 超过开始说的最大数
			return true
		}
		// 4个方向尝试
		for i := 0; i < directNum; i++ {
			ni := []int{pi[0] + dx[i], pi[1] + dy[i]}
			tag := blockPointers.get(ni)
			if tag == targetTag {
				// 找到终点
				return true
			}
			if ni[0] > min && ni[0] < max && ni[1] > min && ni[1] < max && tag == whiteTag {
				// 改为源点
				// 1: 代表该点已经访问过了
				// 2: 反向查找时到了该点一定能到源点
				blockPointers.add(ni, sourceTag)
				queue = append(queue, ni)
			}
		}
	}
	return false
}
