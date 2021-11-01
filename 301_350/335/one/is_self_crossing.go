package one

func IsSelfCrossing(distance []int) bool {
	for i := 3; i < len(distance); i++ {
		if distance[i-1] <= distance[i-3] && distance[i-2] <= distance[i] {
			return true
		}
		if i >= 4 && distance[i-4] < distance[i-2] && distance[i-3] == distance[i-1] && distance[i-2]-distance[i-4] <= distance[i] {
			return true
		}
		if i >= 5 && distance[i-1] <= distance[i-3] && distance[i-4] <= distance[i-2] && distance[i-3]-distance[i-5] <= distance[i-1] && distance[i-2]-distance[i-4] <= distance[i] {
			return true
		}
	}
	return false
}
