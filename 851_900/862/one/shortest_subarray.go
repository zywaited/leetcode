package one

func ShortestSubarray(A []int, K int) int {
	ans := -1
	// 每个数当前的索引和负数之间的和
	q := make([][2]int, 0, len(A))
	// 为了能够判断是否满足
	tk := K
	sum := 0
	ll := 0
	// 开始的节点
	fi := 0
	for i, num := range A {
		sum += num
		if num < 0 {
			// 做下合并处理
			tk += num
			// 如果觉得下面循环耗时间可以多加个数组(实际上整体也就O(N))
			for j := len(q) - 1; j >= 0 && A[q[j][0]] >= 0; j-- {
				num += A[q[len(q)-1][0]]
				q = q[:len(q)-1]
			}
			// 如果实际小于0继续合并到前一个
			for num <= 0 && len(q) > 0 {
				num += q[len(q)-1][1]
				A[i] += A[q[len(q)-1][0]] // 改变负数总和(如果觉得改变原数据不友好，可以保存在q中)
				q = q[:len(q)-1]
			}
		}
		q = append(q, [2]int{i, num})
		ll++
		// tk可能小于0(所以这里判断下长度或sum>0)
		for len(q) > 0 && sum >= tk {
			// 满足条件
			if sum >= K && (ans == -1 || ll < ans) {
				ans = ll
				// 不知道这优化用处大不大
				if ans == 1 {
					return ans
				}
			}
			// 舍弃第一个组合数
			if q[0][1] <= 0 || sum-q[0][1] >= K {
				sum -= q[0][1]
				// 还原tk
				if A[q[0][0]] < 0 {
					tk -= A[q[0][0]]
				}
				ll -= q[0][0] - fi + 1
				fi = q[0][0] + 1
				q = q[1:]
				continue
			}
			// 尝试舍弃第一组第一个数
			if fi < q[0][0] && sum >= K {
				sum -= A[fi]
				q[0][1] -= A[fi]
				ll--
				fi++
				continue
			}
			break
		}
	}
	return ans
}
