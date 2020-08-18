package one

var directs = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	oldColor := image[sr][sc]
	if oldColor == newColor {
		return image
	}
	image[sr][sc] = newColor
	queue := make([]int, 0, len(image)*len(image[0]))
	queue = append(queue, sr*51+sc)
	for len(queue) > 0 {
		r, c := queue[len(queue)-1]/51, queue[len(queue)-1]%51
		queue = queue[:len(queue)-1]
		for _, direct := range directs {
			nr, nc := r+direct[0], c+direct[1]
			if nr < 0 || nc < 0 || nr >= len(image) || nc >= len(image[nr]) {
				continue
			}
			if image[nr][nc] != oldColor {
				continue
			}
			image[nr][nc] = newColor
			queue = append(queue, nr*51+nc)
		}
	}
	return image
}
