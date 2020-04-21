package one

type data struct {
	num  int
	time int
}

type station struct {
	start string
	time  int
}

type UndergroundSystem struct {
	stationTime map[string]map[string]*data
	user        map[int]*station
}

func Constructor() UndergroundSystem {
	return UndergroundSystem{
		stationTime: make(map[string]map[string]*data),
		user:        make(map[int]*station),
	}
}

func (us *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	// 重复利用内存
	if us.user[id] == nil {
		us.user[id] = &station{
			start: stationName,
			time:  t,
		}
		return
	}
	us.user[id].start = stationName
	us.user[id].time = t
}

func (us *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	// 不用管逻辑问题
	if us.stationTime[stationName] == nil {
		us.stationTime[stationName] = make(map[string]*data)
	}
	if us.stationTime[stationName][us.user[id].start] == nil {
		us.stationTime[stationName][us.user[id].start] = &data{}
	}
	us.stationTime[stationName][us.user[id].start].num++
	us.stationTime[stationName][us.user[id].start].time += t - us.user[id].time
}

func (us *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	// 不用管逻辑问题
	return float64(us.stationTime[endStation][startStation].time) / float64(us.stationTime[endStation][startStation].num)
}
