package one

func Evaluate(expression string) int {
	ans, _ := parser(make(map[string]int), []byte(expression), 0)
	return ans
}

func parser(tokens map[string]int, expression []byte, index int) (int, int) {
	if expression[index] == '(' {
		num, nextIndex := parser(tokens, expression, index+1)
		return num, nextIndex + 1 // skip )
	}
	newTokens := copyTokens(tokens)
	token, nextIndex := getVariableName(expression, index)
	nextIndex++ // skip space
	switch token {
	case "let":
		var (
			variable string
			num      int
		)
		for nextIndex < len(expression) {
			if expression[nextIndex] == '(' { // return expr
				return parser(newTokens, expression, nextIndex)
			}
			if isNumber(expression[nextIndex]) { // return number
				return getNumber(expression, nextIndex)
			}
			variable, nextIndex = getVariableName(expression, nextIndex)
			if expression[nextIndex] == ')' {
				return newTokens[variable], nextIndex
			}
			nextIndex++ // skip space
			num, nextIndex = getExprValue(newTokens, expression, nextIndex)
			newTokens[variable] = num
			nextIndex++ // skip space
		}
	case "add":
		first, nextIndex := getExprValue(newTokens, expression, nextIndex)
		second, nextIndex := getExprValue(newTokens, expression, nextIndex+1)
		return first + second, nextIndex
	case "mult":
		first, nextIndex := getExprValue(newTokens, expression, nextIndex)
		second, nextIndex := getExprValue(newTokens, expression, nextIndex+1)
		return first * second, nextIndex
	}
	panic("unreachable")
}

func getExprValue(tokens map[string]int, expression []byte, index int) (int, int) {
	// number
	if isNumber(expression[index]) {
		return getNumber(expression, index)
	}
	// expr
	if expression[index] == '(' {
		return parser(tokens, expression, index)
	}
	// variable
	if isLowLetter(expression[index]) {
		variable, index := getVariableName(expression, index)
		return tokens[variable], index
	}
	panic("unreachable")
}

func getNumber(expression []byte, index int) (int, int) {
	neg := false
	num := 0
	if expression[index] == '-' {
		neg = true
		index++
	}
	for ; expression[index] >= '0' && expression[index] <= '9'; index++ {
		num = num*10 + int(expression[index]-'0')
	}
	if neg {
		num = -num
	}
	return num, index
}

func getVariableName(expression []byte, index int) (string, int) {
	curr := index
	for ; isLowLetter(expression[curr]) || isNumber(expression[curr]); curr++ {
	}
	return string(expression[index:curr]), curr
}

func copyTokens(tokens map[string]int) map[string]int {
	tmpTokens := make(map[string]int, len(tokens))
	for token, num := range tokens {
		tmpTokens[token] = num
	}
	return tmpTokens
}

func isNumber(c byte) bool {
	return c == '-' || (c >= '0' && c <= '9')
}

func isLowLetter(c byte) bool {
	return c >= 'a' && c <= 'z'
}
