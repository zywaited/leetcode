package one

import (
	"math/rand"
	"time"
)

type RandomizedCollection struct {
	indexes map[int]map[int]struct{} // 每个数字对应的索引
	nums    []int                    // rand与数字个数正相关
}

func Constructor() RandomizedCollection {
	rand.Seed(time.Now().UnixNano())
	return RandomizedCollection{indexes: make(map[int]map[int]struct{})}
}

func (rc *RandomizedCollection) Insert(val int) bool {
	indexes := rc.indexes[val]
	if indexes == nil {
		indexes = make(map[int]struct{})
		rc.indexes[val] = indexes
	}
	ok := len(indexes) == 0
	indexes[len(rc.nums)] = struct{}{}
	rc.nums = append(rc.nums, val)
	return ok
}

func (rc *RandomizedCollection) Remove(val int) bool {
	indexes := rc.indexes[val]
	if len(indexes) == 0 {
		return false
	}
	index := 0
	for index = range indexes {
		break
	}
	delete(indexes, index)
	nextVal := rc.nums[len(rc.nums)-1]
	rc.nums[index] = nextVal
	rc.nums = rc.nums[:len(rc.nums)-1]
	if index < len(rc.nums) {
		delete(rc.indexes[nextVal], len(rc.nums))
		rc.indexes[nextVal][index] = struct{}{}
	}
	return true
}

func (rc *RandomizedCollection) GetRandom() int {
	return rc.nums[rand.Intn(len(rc.nums))]
}
