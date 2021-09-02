package one

func CanReach(s string, minJump int, maxJump int) bool {
	stack := make([]bool, len(s))
	stack[0] = true
	si := 0
	dp := true
	for i := 1; i < len(s); i++ {
		if s[i] == '1' {
			dp = false
			continue
		}
		for ; i-si > maxJump || (i-si > minJump && !stack[si]); si++ {
		}
		dp = true
		if i-si < minJump || !stack[si] {
			dp = false
		}
		stack[i] = dp
	}
	return dp
}
