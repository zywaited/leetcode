package one

func Game(guess []int, answer []int) int {
	sum := 0
	for i := range guess {
		if guess[i]^answer[i] == 0 {
			sum++
		}
	}
	return sum
}
