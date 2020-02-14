package one

func CheckIfExist(arr []int) bool {
	mp := make(map[int]byte)
	for _, num := range arr {
		if mp[num*2] == 1 || (num%2 == 0 && mp[num/2] == 1) {
			return true
		}
		mp[num] = 1
	}
	return false
}
