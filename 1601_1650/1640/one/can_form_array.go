package one

func CanFormArray(arr []int, pieces [][]int) bool {
	i := 0
	pm := make(map[int]int)
	for index, piece := range pieces {
		pm[piece[0]] = index + 1
	}
	for i < len(arr) {
		if pm[arr[i]] == 0 {
			return false
		}
		piece := pieces[pm[arr[i]]-1]
		j := 0
		for ; i < len(arr) && j < len(piece); j++ {
			if piece[j] != arr[i] {
				return false
			}
			i++
		}
		if j < len(piece) {
			return false
		}
	}
	return true
}
