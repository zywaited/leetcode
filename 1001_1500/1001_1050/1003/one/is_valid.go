package one

func IsValid(s string) bool {
	var q []byte
	for i := range s {
		q = append(q, s[i])
		if len(q) >= 3 && string(q[len(q)-3:]) == "abc" {
			q = q[:len(q)-3]
		}
	}
	return len(q) == 0
}
