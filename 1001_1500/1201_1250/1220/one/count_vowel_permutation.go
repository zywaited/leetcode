package one

const mod = 1e9 + 7

var rules = map[byte][]byte{
	'a': {'e'},
	'e': {'a', 'i'},
	'i': {'a', 'e', 'o', 'u'},
	'o': {'i', 'u'},
	'u': {'a'},
}

func CountVowelPermutation(n int) int {
	// 总共5中字母
	mp := map[byte]int{
		'a': 1,
		'e': 1,
		'i': 1,
		'o': 1,
		'u': 1,
	}
	sum := 0
	// 减少重复申请内存
	//tmp := make(map[byte]int)
	for i := 2; i <= n; i++ {
		// 用于存储当前增加的数据
		tmp := make(map[byte]int)
		for letter, count := range mp {
			for _, letter = range rules[letter] {
				tmp[letter] = (tmp[letter] + count) % mod
			}
		}
		// 更改原始数据
		mp = tmp
		//for letter, count := range tmp {
		//	mp[letter] = count
		//	tmp[letter] = 0
		//}
	}
	for _, count := range mp {
		sum = (sum + count) % mod
	}
	return sum
}
