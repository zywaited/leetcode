package one

func CanCross(stones []int) bool {
	indexes := map[int]bool{}
	for _, index := range stones {
		indexes[index] = true
	}
	finalIndex := stones[len(stones)-1]
	queue := [][2]int{{0, stones[0]}}
	visited := map[[2]int]bool{}
	var tmp [2]int
	for len(queue) > 0 {
		currentIndex := queue[0]
		queue = queue[1:]
		if currentIndex[1]+currentIndex[0]-1 <= finalIndex && currentIndex[1]+currentIndex[0]+1 >= finalIndex {
			return true
		}
		if currentIndex[0]-1 > 0 && indexes[currentIndex[1]+currentIndex[0]-1] {
			tmp = [2]int{currentIndex[0] - 1, currentIndex[1] + currentIndex[0] - 1}
			if !visited[tmp] {
				queue = append(queue, tmp)
				visited[tmp] = true
			}
		}
		if currentIndex[0] > 0 && indexes[currentIndex[1]+currentIndex[0]] {
			tmp = [2]int{currentIndex[0], currentIndex[1] + currentIndex[0]}
			if !visited[tmp] {
				queue = append(queue, tmp)
				visited[tmp] = true
			}
		}
		if currentIndex[0]+1 > 0 && indexes[currentIndex[1]+currentIndex[0]+1] {
			tmp = [2]int{currentIndex[0] + 1, currentIndex[1] + currentIndex[0] + 1}
			if !visited[tmp] {
				queue = append(queue, tmp)
				visited[tmp] = true
			}
		}
	}
	return false
}
