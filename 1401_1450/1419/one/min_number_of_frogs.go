package one

// 每个字符的前一个字符
var prev = map[byte]byte{
	'r': 'c',
	'o': 'r',
	'a': 'o',
	'k': 'a',
}

const (
	sl = 'c'
	el = 'k'
)

func MinNumberOfFrogs(croakOfFrogs string) int {
	ans := 0
	// 空闲
	ea := 0
	// 已经存在的字符数量
	em := make(map[byte]int, 5)
	for i := range croakOfFrogs {
		em[croakOfFrogs[i]]++
		// 不存在当前字符
		if prev[croakOfFrogs[i]] == 0 && croakOfFrogs[i] != sl {
			return -1
		}
		if croakOfFrogs[i] == sl {
			if ea == 0 {
				ans++
			} else {
				ea--
			}
			continue
		}
		// 前置字符不存在
		if em[prev[croakOfFrogs[i]]] == 0 {
			return -1
		}
		em[prev[croakOfFrogs[i]]]--
		if croakOfFrogs[i] == el {
			ea++
		}
	}
	// 最后应该只剩下k
	for c, n := range em {
		if c != el && n > 0 {
			return -1
		}
	}
	return ans
}
