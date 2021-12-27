package one

import "fmt"

func GetHint(secret string, guess string) string {
	x := 0
	nn := [128]int{}
	for i := range guess {
		if secret[i] == guess[i] {
			x++
			continue
		}
		nn[guess[i]]++
	}
	y := 0
	for i := range secret {
		if secret[i] != guess[i] && nn[secret[i]] > 0 {
			y++
			nn[secret[i]]--
		}
	}
	return fmt.Sprintf("%dA%dB", x, y)
}
