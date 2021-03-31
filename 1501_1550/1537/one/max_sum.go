package one

const mod = 1e9 + 7

func MaxSum(nums1 []int, nums2 []int) int {
	ms := [2]int{}
	n2m := make(map[int]bool)
	for _, n := range nums2 {
		n2m[n] = true
	}
	i, j := -1, -1
	for {
		s1 := [2]int{}
		s2 := [2]int{}
		cn := 0
		for i++; i < len(nums1); i++ {
			s1 = sumN(s1, nums1[i])
			if n2m[nums1[i]] {
				cn = nums1[i]
				break
			}
		}
		for j++; j < len(nums2); j++ {
			s2 = sumN(s2, nums2[j])
			if nums2[j] == cn {
				ms = sum(ms, max(s1, s2))
				break
			}
		}
		if cn == 0 {
			ms = sum(ms, max(s1, s2))
			break
		}
	}
	return ms[1]
}

func sum(n, sn [2]int) [2]int {
	n[0] += sn[0] + (n[1]+sn[1])/mod
	n[1] = (n[1] + sn[1]) % mod
	return n
}

func sumN(n [2]int, sn int) [2]int {
	n[0] += (n[1] + sn) / mod
	n[1] = (n[1] + sn) % mod
	return n
}

func max(f, s [2]int) [2]int {
	if f[0] > s[0] || (f[0] == s[0] && f[1] >= s[1]) {
		return f
	}
	return s
}
