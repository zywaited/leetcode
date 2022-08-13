package one

func MaxChunksToSorted(arr []int) int {
	blocks := make([]int, 0, len(arr))
	for _, num := range arr {
		length := len(blocks)
		for ; length > 0 && num < blocks[length-1]; length-- {
		}
		if length == len(blocks) {
			blocks = append(blocks, num)
			continue
		}
		blocks = append(blocks[:length], blocks[len(blocks)-1])
	}
	return len(blocks)
}
