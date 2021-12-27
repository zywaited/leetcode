package one

/**
 * 与投票法类似，超过n/3最多就两个数
 */

func MajorityElement(nums []int) []int {
	f, s := -1, 0  // 虚拟两个不存在的数
	ff, sf := 0, 0 // 投票数
	for _, num := range nums {
		// 没人投票或者就是当前值
		if (ff == 0 || num == f) && num != s {
			ff++
			f = num
			continue
		}

		if sf == 0 || num == s {
			sf++
			s = num
			continue
		}

		// 两者都不是
		ff--
		sf--
	}

	var r []int
	if ff <= 0 && sf <= 0 {
		return r
	}

	// 选出来的数验证是否真的符合要求
	n := len(nums) / 3
	cf, cs := 0, 0
	for _, num := range nums {
		if num == f {
			cf++
			continue
		}

		if num == s {
			cs++
		}
	}

	if cf > n {
		r = append(r, f)
	}

	if cs > n {
		r = append(r, s)
	}

	return r
}
