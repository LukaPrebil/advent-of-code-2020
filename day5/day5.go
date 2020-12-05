package day5

import (
	"log"
	"sort"
)

func GetMaxSeatID() int {
	maxId := 0
	for _, seat := range parseFile() {
		id := seat.ComputeID()
		if id > maxId {
			maxId = id
		}
	}
	return maxId
}

func GetMySeatID() int {
	var previous, current, next int
	ids := getAllIdsSorted()
	for i := 1; i < len(ids)-1; i++ {
		previous, current, next = ids[i-1], ids[i], ids[i+1]
		if previous != current - 1 || current + 1 != next {
			return next - 1
		}
	}
	log.Panic("Didnt find id")
	return -1
}

func getAllIdsSorted() []int {
	ids := make([]int, 0)
	for _, seat := range parseFile() {
		ids = append(ids, seat.ComputeID())
	}
	sort.Ints(ids)
	return ids
}

type IDComputer interface {
	ComputeID() int
}

func (s seat) ComputeID() int {
	return s.row * 8 + s.column
}