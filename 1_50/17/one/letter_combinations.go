package one

var dict = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func bc(i int, digits, s string, resp *[]string) {
	if i == len(digits) {
		*resp = append(*resp, s)
		return
	}

	d, ok := dict[string(digits[i])]
	if !ok {
		return
	}
	for j := 0; j < len(d); j++ {
		bc(i+1, digits, s+string(d[j]), resp)
	}
}

func LetterCombinations(digits string) []string {
	var resp []string
	if len(digits) < 1 {
		return resp
	}
	bc(0, digits, "", &resp)
	return resp
}
