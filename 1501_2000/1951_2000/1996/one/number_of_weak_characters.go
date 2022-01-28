package one

import "sort"

func NumberOfWeakCharacters(properties [][]int) int {
	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] == properties[j][0] {
			return properties[j][1] < properties[i][1]
		}
		return properties[i][0] < properties[j][0]
	})
	ans := 0
	i := len(properties) - 1
	maxProperty := i
	for i--; i >= 0; i-- {
		if properties[i][0] < properties[maxProperty][0] && properties[i][1] < properties[maxProperty][1] {
			ans++
			continue
		}
		if properties[maxProperty][1] < properties[i][1] {
			maxProperty = i
		}
	}
	return ans
}
