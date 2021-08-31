package two

func CorpFlightBookings(bookings [][]int, n int) []int {
	seats := make([]int, n)
	for _, booking := range bookings {
		seats[booking[0]-1] += booking[2]
		if booking[1] < n {
			seats[booking[1]] -= booking[2]
		}
	}
	for i := 1; i < n; i++ {
		seats[i] += seats[i-1]
	}
	return seats
}
