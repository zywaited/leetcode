package one

func Makesquare(matchsticks []int) bool {
	sum := 0
	max := 0
	for _, ml := range matchsticks {
		sum += ml
		if max < ml {
			max = ml
		}
	}
	if sum%4 > 0 {
		return false
	}
	side := sum / 4
	if side < max {
		return false
	}
	valid := make([]byte, 1<<uint(len(matchsticks)))
	var bc func(curr, length int) byte
	bc = func(curr, length int) byte {
		if length == sum && curr+1 == len(valid) {
			valid[curr] = 1
		}
		if valid[curr] != 0 {
			return valid[curr]
		}
		valid[curr] = 2
		ml := side - length%side
		for i := 0; i < len(matchsticks) && valid[curr] != 1; i++ {
			if curr&(1<<uint(i)) == 0 && matchsticks[i] <= ml {
				valid[curr] = bc(curr|(1<<uint(i)), length+matchsticks[i])
			}
		}
		return valid[curr]
	}
	return bc(0, 0) == 1
}
