package one

func FindOcurrences(text string, first string, second string) []string {
	ans := make([]string, 0)
	i := 0
	n := 0
	for j := 0; j <= len(text); j++ {
		if j == len(text) || text[j] == ' ' {
			s := text[i:j]
			i = j + 1
			if n == 0 && s == first {
				n++
				continue
			}
			if n == 1 && s == second {
				n++
				continue
			}
			if n == 2 {
				ans = append(ans, s)
			}
			n = 0
			if s != first {
				continue
			}
			n++
			if s == second {
				n++
			}
		}
	}
	return ans
}
