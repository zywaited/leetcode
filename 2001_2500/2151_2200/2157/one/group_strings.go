package one

func GroupStrings(words []string) []int {
	indexes := make([]int, len(words))
	for index := range indexes {
		indexes[index] = index
	}
	type mapKey struct {
		nums    [26]int
		keyType byte // 0: 原始数据 1：可替换
	}
	mapIndexes := make(map[mapKey]int)
	for index, word := range words {
		key := mapKey{}
		for i := range word {
			key.nums[word[i]-'a']++
		}
		// 先查找有没有相同的
		if nextIndex, ok := mapIndexes[key]; ok {
			union(indexes, index, nextIndex)
		}
		mapIndexes[key] = index
		for i := 0; i < 26; i++ {
			key.nums[i]++
			key.keyType = 0
			// 新增一个字符看是否有存在的
			if nextIndex, ok := mapIndexes[key]; ok {
				union(indexes, index, nextIndex)
			}
			key.nums[i]--
			if key.nums[i] > 0 {
				key.nums[i]--
				// 删除一个字符看是否有存在的
				if nextIndex, ok := mapIndexes[key]; ok {
					union(indexes, index, nextIndex)
				}
				// 是否有存在可以替换的
				key.keyType = 1
				if nextIndex, ok := mapIndexes[key]; ok {
					union(indexes, index, nextIndex)
				}
				mapIndexes[key] = index
				key.nums[i]++
			}
		}
	}
	group := 0
	nums := make([]int, len(words))
	maxNum := 0
	for index := range words {
		nextIndex := find(indexes, index)
		if nextIndex == index {
			group++
		}
		nums[nextIndex]++
		if maxNum < nums[nextIndex] {
			maxNum = nums[nextIndex]
		}
	}
	return []int{group, maxNum}
}

func find(indexes []int, index int) int {
	for indexes[index] != index {
		indexes[index] = indexes[indexes[index]]
		index = indexes[index]
	}
	return index
}

func union(indexes []int, firstIndex, secondIndex int) {
	indexes[find(indexes, firstIndex)] = indexes[find(indexes, secondIndex)]
}
