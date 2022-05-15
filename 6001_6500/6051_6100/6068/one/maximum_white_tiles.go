package one

import "sort"

func MaximumWhiteTiles(tiles [][]int, carpetLen int) int {
	sort.Slice(tiles, func(i, j int) bool {
		return tiles[i][0] < tiles[j][0]
	})
	ans := 0
	i := 0
	j := 0
	l := 0
	for ; j < len(tiles) && ans < carpetLen; j++ {
		l += tiles[j][1] - tiles[j][0] + 1
		for tiles[i][1]+carpetLen <= tiles[j][1] {
			l -= tiles[i][1] - tiles[i][0] + 1
			i++
		}
		cl := l - (tiles[i][1] - tiles[i][0] + 1)
		dl := carpetLen - (tiles[j][1] - tiles[i][1])
		if tiles[i][1]-dl+1 < tiles[i][0] {
			dl = tiles[i][1] - tiles[i][0] + 1
		}
		ans = max(ans, cl+dl)
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}
