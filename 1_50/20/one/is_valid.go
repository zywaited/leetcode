package one

func IsValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := range s {
		if s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i])
			continue
		}
		sl := len(stack)
		if sl == 0 {
			return false
		}
		prev := stack[sl-1]
		if s[i] == ')' && prev != '(' {
			return false
		}
		if s[i] == ']' && prev != '[' {
			return false
		}
		if s[i] == '}' && prev != '{' {
			return false
		}
		stack = stack[:sl-1]
	}
	return len(stack) == 0
}
