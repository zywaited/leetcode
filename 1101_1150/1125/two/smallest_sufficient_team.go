package two

func SmallestSufficientTeam(skills []string, people [][]string) []int {
	fs := 1<<uint(len(skills)) - 1
	// 技能点位置
	skillIndex := make(map[string]uint)
	for i, skill := range skills {
		skillIndex[skill] = uint(i)
	}

	// 本身原始回溯是可以完成本题的，不过效率太低，偶尔执行超时
	// 优化点:
	// 1: 去除技能可被替代的同学
	// 2: 找出独有技能的同学

	// 技能拥有人数
	sp := make(map[string][]int)

	// 先计算每个同学的技能点
	pl := len(people)
	ps := make([]int, pl)
	for i, ss := range people {
		for _, skill := range ss {
			if _, ok := skillIndex[skill]; !ok {
				continue
			}
			ps[i] |= 1 << skillIndex[skill]
			sp[skill] = append(sp[skill], i)
		}
	}

	var (
		bc  func(int)
		tmp []int
		r   []int
	)

	// 替代人员
	for i := 0; i < pl; i++ {
		for j := i + 1; j < pl; j++ {
			if ps[i]|ps[j] == ps[i] {
				// i 包含 j
				ps[j] = 0
				continue
			}
			if ps[i]|ps[j] == ps[j] {
				// j 包含 i
				ps[i] = 0
				break
			}
		}
	}
	// 添加单独的技能人员
	cs := 0
	for _, p := range sp {
		// 可能一个同学拥有多个独有技能，所以去重
		if len(p) > 1 || ps[p[0]] == 0 {
			continue
		}
		tmp = append(tmp, p[0])
		cs |= ps[p[0]]
		ps[p[0]] = 0
	}

	bc = func(skill int) {
		// 是否团队人数最小
		if len(r) != 0 && len(tmp) >= len(r) {
			return
		}
		// 是否已经找到
		if skill == fs {
			r = append([]int{}, tmp...)
			return
		}
		// 当前已经拥有的技能数
		num := uint(0)
		for (skill>>num)&1 == 1 {
			num++
		}
		// 继续找下一位成员
		for i, s := range ps {
			// 找到下一个技能拥有者
			if (s>>num)&1 != 1 {
				continue
			}
			// 加入新成员
			tmp = append(tmp, i)
			bc(skill | s)
			// 去除该成员
			tmp = tmp[:len(tmp)-1]
		}
	}
	bc(cs)
	return r
}
