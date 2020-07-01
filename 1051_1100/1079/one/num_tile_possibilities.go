package one

func NumTilePossibilities(tiles string) int {
	cn := make([]int, 26)
	for ti := range tiles {
		cn[tiles[ti]-'A']++
	}
	return bc(cn)
}

func bc(cn []int) int {
	num := 0
	for c := range cn {
		if cn[c] == 0 {
			continue
		}
		num++
		cn[c]--
		num += bc(cn)
		cn[c]++
	}
	return num
}
