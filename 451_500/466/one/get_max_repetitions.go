package one

func GetMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	// 只能找循环体(题目要求有顺序)
	// 找到N个s1组成S2
	// 找到循环体时s1,s2用的数量
	sc1 := 0
	sc2 := 0
	// [2]int 保存了前一次sc1,sc2使用数量(方便后面计算)
	indexes := make(map[int][2]int, len(s2))
	indexes[0] = [2]int{sc1, sc2}
	// s2到了哪里
	index := 0
	// 组成循环时前面已经找到的数量
	//循环所需的数量
	pn := [2]int{}
	ln := [2]int{}
	for {
		sc1++
		for i := range s1 {
			if s1[i] == s2[index] {
				index++
				if index == len(s2) {
					// s2匹配完了
					index = 0
					sc2++
				}
			}
		}
		if sc1 == n1 || (index == 0 && sc2 == 0) {
			// 这里所有数据已被使用，那么计算数量(或者根本没有能够匹配的数据)
			return sc2 / n2
		}
		// 如果没有使用完，那么是否循环了
		if _, ok := indexes[index]; !ok {
			indexes[index] = [2]int{sc1, sc2}
			continue
		}
		// 这时候就要计算多个数量可以组成循环
		pn = indexes[index]
		ln[0], ln[1] = sc1-indexes[index][0], sc2-indexes[index][1]
		break
	}
	// s2的总数量
	ans := pn[1] + (n1-pn[0])/ln[0]*ln[1]
	// 循环体结束后可能还有数据，判断能不能组合
	tsn1 := (n1 - pn[0]) % ln[0]
	for ; tsn1 > 0; tsn1-- {
		for i := range s1 {
			if s1[i] == s2[index] {
				index++
				if index == len(s2) {
					ans++
					index = 0
				}
			}
		}
	}
	return ans / n2
}
