package two

func RemoveStones(stones [][]int) int {
	indexes := make(map[int]int)
	find := func(index int) int {
		if _, ok := indexes[index]; !ok {
			indexes[index] = index
		}
		for indexes[index] != index {
			indexes[index] = indexes[indexes[index]]
			index = indexes[index]
		}
		return index
	}
	union := func(firstIndex, secondIndex int) {
		indexes[find(firstIndex)] = indexes[find(secondIndex)]
	}
	for _, stone := range stones {
		union(stone[0], stone[1]+10001)
	}
	cnt := 0
	for index := range indexes {
		if index == find(index) {
			// 只剩一个点
			cnt++
		}
	}
	return len(stones) - cnt
}
