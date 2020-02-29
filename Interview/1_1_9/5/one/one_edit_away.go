package one

func OneEditAway(first string, second string) bool {
	// 因为只能进行一次编辑, 所以两个字符串的长度差在1以内
	diff := len(first) - len(second)
	if diff > 1 || diff < -1 {
		return false
	}
	num := 0
	i, j := 0, 0
	for i < len(first) && j < len(second) {
		if first[i] == second[j] {
			i++
			j++
			continue
		}
		// 已经不能更改
		if num > 0 {
			return false
		}
		num++
		// 多一个刚好删除
		if diff == 1 {
			i++
			continue
		}
		if diff == -1 {
			j++
			continue
		}
		// 相等就只能替换了
		i++
		j++
	}
	return true
}
