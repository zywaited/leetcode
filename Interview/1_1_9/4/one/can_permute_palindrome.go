package one

func CanPermutePalindrome(s string) bool {
	np := make(map[byte]int)
	for i := range s {
		np[s[i]]++
	}
	num := 0
	for _, n := range np {
		num += n % 2
	}
	return num == len(s)%2
}
