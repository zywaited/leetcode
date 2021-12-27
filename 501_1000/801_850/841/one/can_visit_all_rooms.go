package one

func CanVisitAllRooms(rooms [][]int) bool {
	// 每个房间是否被打开
	rm := make([]bool, len(rooms))
	// 打开的房间数
	rn := 0

	queue := make([]int, 0, len(rooms))
	add := func(index int) {
		for _, nr := range rooms[index] {
			if !rm[nr] {
				queue = append(queue, nr)
				rm[nr] = true
				rn++
			}
		}
	}
	adds := func(index int) {
		if !rm[index] {
			rm[index] = true
			rn++
			add(index)
		}
	}
	adds(0)
	for len(queue) > 0 {
		r := queue[0]
		queue = queue[1:]
		add(r)
		if rn == len(rooms) {
			break
		}
	}

	return rn == len(rooms)
}
