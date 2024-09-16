package one

func maximumRobots(chargeTimes []int, runningCosts []int, budget int64) int {
	res, n := 0, len(chargeTimes)
	runningCostSum := int64(0)
	var q []int
	for i, j := 0, 0; i < n; i++ {
		runningCostSum += int64(runningCosts[i])
		for len(q) > 0 && chargeTimes[q[len(q)-1]] <= chargeTimes[i] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
		for j <= i && int64(i-j+1)*runningCostSum+int64(chargeTimes[q[0]]) > budget {
			if len(q) > 0 && q[0] == j {
				q = q[1:]
			}
			runningCostSum -= int64(runningCosts[j])
			j++
		}
		res = max(res, i-j+1)
	}
	return res
}
