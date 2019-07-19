package three

func SmallestSufficientTeam(skills []string, people [][]string) []int {
	// 技能点位置
	skillIndex := make(map[string]uint)
	for i, skill := range skills {
		skillIndex[skill] = uint(i)
	}

	// 跟回溯一样，两方面可以提前优化好，不过没有相比优化后的回溯也不会太差

	// 先计算每个同学的技能点
	pl := len(people)
	ps := make([]int, pl)
	for i, ss := range people {
		for _, skill := range ss {
			if _, ok := skillIndex[skill]; !ok {
				continue
			}
			ps[i] |= 1 << skillIndex[skill]
		}
	}

	// 记住中间状态
	mp := make(map[int][]int)
	// 所有技能
	fs := (1 << uint(len(skills))) - 1
	var dp func(int, int, int) []int
	dp = func(i, now, need int) []int {
		if _, ok := mp[need]; ok {
			return mp[need]
		}

		for j := i + 1; j < pl; j++ {
			// 无技能或者被包含
			if ps[j] == 0 || now != 0 && ps[j]|now == now {
				continue
			}
			// 一个人
			if ps[j]|need == ps[j] {
				mp[need] = []int{j}
				break
			}
			tmp := dp(j, now|ps[j], (now|ps[j])^fs)
			// 找不到对应人员或者人数比现在还多
			if len(tmp) == 0 || len(mp[need]) != 0 && len(tmp) >= len(mp[need])-1 {
				continue
			}
			mp[need] = append([]int{j}, tmp...)
		}
		// 没找到则置空
		if len(mp[need]) == 0 {
			mp[need] = make([]int, 0)
		}
		return mp[need]
	}
	return dp(-1, 0, fs)
}
