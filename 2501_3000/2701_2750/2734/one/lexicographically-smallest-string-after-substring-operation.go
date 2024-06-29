package one

func smallestString(s string) string {
	newBytes := []byte(s)
	changed := false
	for i := range newBytes {
		if newBytes[i] != 'a' {
			newBytes[i] -= 1
			changed = true
			continue
		}
		if changed {
			break
		}
	}
	if !changed {
		newBytes[len(newBytes)-1] = 'z'
	}
	return string(newBytes)
}
