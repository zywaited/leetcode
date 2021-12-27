package one

// 拆分逻辑
// 假设从nums1中取i个数字，则需要从nums2中取k-i个数字，然后对返回的各种情况比对取出最大值
func MaxNumber(nums1, nums2 []int, k int) []int {
	rs := make([]int, 0)
	fl := len(nums1)
	sl := len(nums2)
	maxK := min(fl, k)   // 可以从nums1中取出的最大个数
	minK := max(0, k-sl) // 可以从nums1中取出的最小个数
	for ; minK <= maxK; minK++ {
		fns, sns := maxKNumber(nums1, minK), maxKNumber(nums2, k-minK)
		tms := make([]int, 0, k)
		// 从fns, sns找最大值，类似maxKNumber
		for len(fns) > 0 || len(sns) > 0 {
			if compare(fns, sns) {
				tms = append(tms, fns[0])
				fns = fns[1:]
				continue
			}
			tms = append(tms, sns[0])
			sns = sns[1:]
		}
		// 最后与rs做对比
		if compare(tms, rs) {
			rs = tms
		}
	}
	return rs
}

func maxKNumber(nums []int, k int) []int {
	if k > 0 {
		rs := make([]int, 0, k)
		rl := 0
		nl := len(nums) - k // 可以去除的个数
		for _, num := range nums {
			for nl > 0 && rl > 0 && rs[rl-1] < num {
				// 找到最大值
				rs = rs[:rl-1]
				nl--
				rl--
			}
			if rl < k {
				rs = append(rs, num)
				rl++
				continue
			}
			nl-- // 舍弃后可用数字减少
		}
		return rs
	}
	return nil
}

// 最高位大小
func compare(nums1, nums2 []int) bool {
	fl := len(nums1)
	sl := len(nums2)
	m := min(fl, sl)
	for i := 0; i < m; i++ {
		if nums1[i] > nums2[i] {
			return true
		}
		if nums1[i] < nums2[i] {
			return false
		}
	}
	return fl >= sl
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}
