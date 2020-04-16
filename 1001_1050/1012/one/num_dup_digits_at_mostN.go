package one

func NumDupDigitsAtMostN(N int) int {
	// 没有任何重复的数量
	dn := 0
	// 阶乘值
	fn := 1
	tn := N
	nn := make([]int, 0)
	cn := 0
	for tn > 0 {
		cn = tn % 10
		tn /= 10
		nn = append(nn, cn)
		if tn == 0 {
			break
		}
		// 前面还有一位(阶乘)
		dn += fn * 9
		fn *= 10 - len(nn)
	}
	// 最后不是全变化的计算
	ff := 10
	fn *= ff
	en := 0 // 已经存在的数据
	for i := len(nn) - 1; i >= 0; i-- {
		fn /= ff
		ff--
		if i == len(nn)-1 {
			dn += fn * (nn[i] - 1)
			en |= 1 << uint(nn[i])
			continue
		}
		// 先判断是否是大于0
		if nn[i] > 0 {
			// 现在已经存在的数字个数(1的个数)
			cn = nn[i] - bitCount(en&(1<<uint(nn[i])-1))
			// 一定会重复(只能保持原数据)
			if cn > 0 {
				dn += fn * cn
			}
		}
		// 如果重复
		if en&(1<<uint(nn[i])) > 0 {
			return N - dn
		}
		en |= 1 << uint(nn[i])
	}
	// 本身也算一个不重复的
	return N - dn - 1
}

func bitCount(f int) int {
	ans := 0
	for f > 0 {
		ans++
		f &= f - 1
	}
	return ans
}
