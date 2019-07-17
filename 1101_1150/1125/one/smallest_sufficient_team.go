package one

func SmallestSufficientTeam(skills []string, people [][]string) []int {
	length := 1 << uint(len(skills))
	// 所有可能的团队技能点
	dps := make([][]int, length)
	// 需要的技能点索引
	skillIndex := make(map[string]int)
	for i := range skills {
		skillIndex[skills[i]] = i
	}
	for i, pe := range people {
		// 当前同学的技能点
		x := 0
		for _, skill := range pe {
			x |= 1 << uint(skillIndex[skill])
		}
		// 该同学没有技能点
		if x == 0 {
			continue
		}
		for k, dp := range dps {
			if k != 0 && len(dp) == 0 {
				continue
			}
			// 团队技能点
			skill := x | k
			// 该新技能点人员数量对比
			if len(dps[skill]) != 0 && len(dps[skill]) <= len(dp)+1 {
				continue
			}
			// 支撑新技能的人员，也就是当前技能团队加上新成员
			dps[skill] = append([]int{i}, dp...)
		}
	}
	// 返回所有技能都有的团队
	return dps[length-1]
}
