package one

func CountCharacters(words []string, chars string) int {
	actual := make(map[byte]int, len(chars))
	for i := range chars {
		actual[chars[i]]++
	}
	ans := 0
	// 不重复初始化
	excepted := make(map[byte]int)
	num := 0
	flag := true
	for _, word := range words {
		flag = true
		for i := range word {
			if excepted[word[i]] < num {
				excepted[word[i]] = num
			}
			excepted[word[i]]++
			if excepted[word[i]]-num > actual[word[i]] {
				flag = false
				break
			}
		}
		if flag {
			ans += len(word)
		}
		num += len(word)
	}
	return ans
}
