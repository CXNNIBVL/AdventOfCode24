package cmd

import (
	"math/rand"
	"sort"
)

func makeList(numItems uint, bound uint) []uint {
	list := make([]uint, numItems)

	for i := uint(0); i < numItems; i++ {
		list[i] = uint(rand.Intn(int(bound)))
	}

	return list
}

func sortList(list []uint) {
	sort.Slice(list, func(i, j int) bool {
		return list[i] < list[j]
	})
}

func GetTotalDistance(leftList, rightList []uint) uint {
	sortList(leftList)
	sortList(rightList)

	var totalDistance uint = 0

	for ix := 0; ix < len(leftList); ix++ {
		var diff, r, l uint = 0, rightList[ix], leftList[ix]

		if r < l {
			diff = l - r
		} else {
			diff = r - l
		}

		totalDistance = totalDistance + diff
	}

	return totalDistance
}
